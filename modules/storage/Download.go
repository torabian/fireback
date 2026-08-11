// Endpoints for reading back uploads that were stored via the tus protocol
// handler in handler.go/store.go. These are mounted on their own base path
// (not nested under the tus basePath) since tusd's routed handler already
// claims a catch-all "*id" route under its own mount point and gin's router
// will not let a more specific route coexist with a catch-all at the same
// path.
package storage

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// MountDownloads wires read-only endpoints for stored uploads onto r, e.g.
// MountDownloads(r, "/downloads", store, cfg).
//
//	GET  basePath/:id      - JSON metadata (size, filename, completion state, ...)
//	HEAD basePath/:id/raw  - same headers a GET would send, without a body
//	GET  basePath/:id/raw  - the file bytes; honors Range/If-Range/If-None-Match
//	                         so download managers can resume or split a
//	                         download across several parallel connections
//
// Like Mount, every request runs through cfg.Authenticate; when it's set,
// only the upload's recorded owner (or anyone, if it has none recorded) may
// read it - see authorizeOwner.
func MountDownloads(r gin.IRouter, basePath string, store Store, cfg *StorageModuleConfig) {
	cfg = cfg.withDefaults()
	basePath = strings.TrimSuffix(basePath, "/")

	r.GET(basePath+"/:id", fileInfoHandler(store, cfg))
	r.HEAD(basePath+"/:id/raw", downloadHandler(store, cfg))
	r.GET(basePath+"/:id/raw", downloadHandler(store, cfg))
}

func fileInfoHandler(store Store, cfg *StorageModuleConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		rec, err := GetFile(c.Request.Context(), store, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rec == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		if cfg.Authenticate != nil {
			auth, err := resolveAuth(c, cfg)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			if !authorizeOwner(auth, rec.UserId) {
				c.JSON(http.StatusForbidden, gin.H{"error": "not your upload"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"id":          rec.ID,
			"size":        rec.Size,
			"offset":      rec.Offset,
			"completed":   rec.Completed,
			"metadata":    rec.MetaData,
			"createdAt":   rec.CreatedAt,
			"updatedAt":   rec.UpdatedAt,
			"completedAt": rec.CompletedAt,
			"claimed":     rec.ClaimedBy != nil,
			"claimedBy":   rec.ClaimedBy,
			"claimedAt":   rec.ClaimedAt,
		})
	}
}

var mimeTypeRe = regexp.MustCompile(`^[a-z-]+/[a-zA-Z0-9.+_-]+$`)

// inlineSafeContentType reports whether contentType is a media type that's
// safe to render directly in the browser (e.g. inside a <video>/<audio>/<img>
// tag) rather than forced through a download prompt. It excludes
// image/svg+xml since an SVG can embed <script>, which would defeat the same
// stored-XSS protection contentTypeAndDisposition is guarding.
func inlineSafeContentType(contentType string) bool {
	switch {
	case strings.HasPrefix(contentType, "video/"), strings.HasPrefix(contentType, "audio/"):
		return true
	case strings.HasPrefix(contentType, "image/"):
		return contentType != "image/svg+xml"
	default:
		return false
	}
}

// contentTypeAndDisposition mirrors tusd's own filterContentType: an
// unrecognized or missing "filetype" always falls back to
// application/octet-stream + attachment, so a maliciously crafted metadata
// value (e.g. text/html) can't turn this into a stored-XSS vector when the
// link is opened directly in a browser. Recognized media types that are safe
// to render inline (video/audio/image, minus SVG) get "inline" instead, so
// e.g. a <video src=".../raw"> can play the file instead of downloading it.
func contentTypeAndDisposition(meta map[string]string) (contentType, disposition string) {
	contentType = "application/octet-stream"
	if filetype := meta["filetype"]; mimeTypeRe.MatchString(filetype) {
		contentType = filetype
	}

	disposition = "attachment"
	if inlineSafeContentType(contentType) {
		disposition = "inline"
	}
	if filename, ok := meta["filename"]; ok {
		disposition += "; filename=" + strconv.Quote(filename)
	}
	return contentType, disposition
}

func downloadHandler(store Store, cfg *StorageModuleConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")

		rec, err := GetFile(ctx, store, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rec == nil || !rec.Completed {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		if cfg.Authenticate != nil {
			auth, err := resolveAuth(c, cfg)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			if !authorizeOwner(auth, rec.UserId) {
				c.JSON(http.StatusForbidden, gin.H{"error": "not your upload"})
				return
			}
		}

		size := rec.Size
		contentType, disposition := contentTypeAndDisposition(rec.MetaData)
		etag := fmt.Sprintf(`"%s-%d"`, rec.ID, size)
		lastModified := rec.UpdatedAt
		if rec.CompletedAt != nil {
			lastModified = *rec.CompletedAt
		}

		w := c.Writer
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Disposition", disposition)
		w.Header().Set("ETag", etag)
		w.Header().Set("Last-Modified", lastModified.UTC().Format(http.TimeFormat))

		if match := c.GetHeader("If-None-Match"); match != "" && match == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		start, end, status, ok := resolveRange(c, size, etag)
		if !ok {
			w.Header().Set("Content-Range", fmt.Sprintf("bytes */%d", size))
			w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
			return
		}

		length := end - start + 1
		w.Header().Set("Content-Length", strconv.FormatInt(length, 10))
		if status == http.StatusPartialContent {
			w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, size))
		}
		w.WriteHeader(status)

		if c.Request.Method == http.MethodHead || length == 0 {
			return
		}

		src, err := store.OpenRange(ctx, id, start)
		if err != nil {
			// Headers and status are already on the wire at this point;
			// there's nothing left to do but stop writing the body.
			return
		}
		defer src.Close()

		io.CopyN(w, src, length)
	}
}

// resolveRange inspects the request's Range/If-Range headers and returns the
// inclusive byte range to serve plus the status code for it (200 for the
// whole file, 206 for a satisfied single range). ok is false only when a
// Range header was present and unsatisfiable, in which case the caller must
// respond 416.
//
// A request with several comma-separated ranges (multipart/byteranges) falls
// back to sending the whole file with 200: download managers that split a
// file across multiple connections do so with one single-range request per
// connection, so multi-range in a single request is an exotic case the RFC
// explicitly allows a server to decline.
func resolveRange(c *gin.Context, size int64, etag string) (start, end int64, status int, ok bool) {
	rangeHeader := c.GetHeader("Range")
	if rangeHeader == "" {
		return 0, size - 1, http.StatusOK, true
	}

	if ifRange := c.GetHeader("If-Range"); ifRange != "" && ifRange != etag {
		return 0, size - 1, http.StatusOK, true
	}

	ranges, err := parseByteRanges(rangeHeader, size)
	if err != nil {
		return 0, 0, 0, false
	}
	if len(ranges) != 1 {
		return 0, size - 1, http.StatusOK, true
	}

	return ranges[0].start, ranges[0].end, http.StatusPartialContent, true
}

type byteRange struct{ start, end int64 } // both inclusive

// parseByteRanges parses an RFC 7233 "Range: bytes=..." header value against
// a resource of the given size, e.g. "bytes=0-499", "bytes=500-" (from byte
// 500 to the end, what a resumed download sends as the offset) or
// "bytes=-500" (the last 500 bytes).
func parseByteRanges(header string, size int64) ([]byteRange, error) {
	const prefix = "bytes="
	if !strings.HasPrefix(header, prefix) {
		return nil, fmt.Errorf("unsupported range unit")
	}

	var ranges []byteRange
	for part := range strings.SplitSeq(strings.TrimPrefix(header, prefix), ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		before, after, found := strings.Cut(part, "-")
		if !found {
			return nil, fmt.Errorf("invalid range %q", part)
		}
		startStr, endStr := strings.TrimSpace(before), strings.TrimSpace(after)

		var r byteRange
		switch {
		case startStr == "":
			n, err := strconv.ParseInt(endStr, 10, 64)
			if err != nil || n <= 0 {
				return nil, fmt.Errorf("invalid suffix range %q", part)
			}
			if n > size {
				n = size
			}
			r = byteRange{start: size - n, end: size - 1}
		case endStr == "":
			start, err := strconv.ParseInt(startStr, 10, 64)
			if err != nil || start < 0 {
				return nil, fmt.Errorf("invalid range %q", part)
			}
			r = byteRange{start: start, end: size - 1}
		default:
			start, err1 := strconv.ParseInt(startStr, 10, 64)
			end, err2 := strconv.ParseInt(endStr, 10, 64)
			if err1 != nil || err2 != nil || start > end {
				return nil, fmt.Errorf("invalid range %q", part)
			}
			if end > size-1 {
				end = size - 1
			}
			r = byteRange{start: start, end: end}
		}

		if r.start < 0 || r.start > r.end || r.start >= size {
			return nil, fmt.Errorf("unsatisfiable range %q", part)
		}
		ranges = append(ranges, r)
	}

	if len(ranges) == 0 {
		return nil, fmt.Errorf("empty range header")
	}
	return ranges, nil
}
