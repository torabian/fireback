package abac

import (
	"fmt"
	"time"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

// Analytics has no entity of its own (it's a pure read-only aggregate over other
// entities) - just the one permission a CRUD set would normally call Query, so root
// (or a future non-root role explicitly granted this) can be gated on viewing it.
var analyticsPerms = NewCrudPermissionSet("root.modules", "analytics", "analytics")
var PERM_ROOT_ANALYTICS_QUERY = analyticsPerms.Query
var ALL_ANALYTICS_PERMISSIONS = analyticsPerms.All

// analyticsCollector mirrors modules/internalstats/Collector.go's own collector -
// one add() call per stat tile, in the exact order they should display, grouped by
// category (see AnalyticsOverviewAction.go's out.fields for the shape this builds).
type analyticsCollector struct {
	items  []abacdefs.AnalyticsOverviewActionResItems
	series []abacdefs.AnalyticsOverviewActionResSeries
}

func (c *analyticsCollector) add(key, label, category, value string, rawValue float64, unit, severity string) {
	c.items = append(c.items, abacdefs.AnalyticsOverviewActionResItems{
		Key:      key,
		Label:    label,
		Category: category,
		Value:    value,
		RawValue: rawValue,
		Unit:     unit,
		Severity: severity,
	})
}

func (c *analyticsCollector) addSeries(key, label, category, chartType, unit string, points []abacdefs.AnalyticsOverviewActionResSeriesPoints) {
	c.series = append(c.series, abacdefs.AnalyticsOverviewActionResSeries{
		Key:       key,
		Label:     label,
		Category:  category,
		ChartType: chartType,
		Unit:      unit,
		Points:    emigo.ArrayReplace(points),
	})
}

// countRateSeverity is percentSeverity's mirror image for AnalyticsOverviewAction's
// own percent-based stats (confirmation rate, 2FA adoption, notification read rate) -
// unlike Collector.go's percentSeverity (used for resource *usage*, where high is
// bad), a low rate here is the problem, so the thresholds run the other direction.
func countRateSeverity(pct float64) string {
	switch {
	case pct < 40:
		return SeverityCriticalAnalytics
	case pct < 75:
		return SeverityWarnAnalytics
	default:
		return SeverityOkAnalytics
	}
}

// Deliberately distinct constants from modules/internalstats' Severity* (unexported
// there) - same four values (ok/warn/critical/info), just spelled out locally so this
// file has no import-time dependency on that module.
const (
	SeverityOkAnalytics       = "ok"
	SeverityWarnAnalytics     = "warn"
	SeverityCriticalAnalytics = "critical"
	SeverityInfoAnalytics     = "info"
)

// AnalyticsOverviewAction computes every headline stat and chart dataset fresh, on
// every call - this is a root-only admin screen (see AllowOnRoot below), not a
// high-traffic public endpoint, so there's no caching layer here; add one (a
// periodic snapshot table, a short-lived in-memory cache) if this ever needs to
// serve a dashboard that auto-refreshes on a tight interval against a large database.
func AnalyticsOverviewAction(c abacdefs.AnalyticsOverviewActionRequest) (*abacdefs.AnalyticsOverviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_ANALYTICS_QUERY},
		AllowOnRoot:    true,
	}); err != nil {
		return nil, err
	}

	tx := fireback.GetDbRef()
	now := time.Now().UTC()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	last7Days := now.AddDate(0, 0, -7)
	last30Days := now.AddDate(0, 0, -30)

	col := &analyticsCollector{}

	// --- Users ---------------------------------------------------------------
	usersTotal := countRows(tx, &abacdefs.UserEntity{}, "")
	usersToday := countRows(tx, &abacdefs.UserEntity{}, "created_at >= ?", todayStart)
	usersLast7 := countRows(tx, &abacdefs.UserEntity{}, "created_at >= ?", last7Days)
	usersLast30 := countRows(tx, &abacdefs.UserEntity{}, "created_at >= ?", last30Days)

	col.add("users.total", "Total Users", "Users", fmt.Sprintf("%d", usersTotal), float64(usersTotal), "count", SeverityInfoAnalytics)
	col.add("users.newToday", "New Users Today", "Users", fmt.Sprintf("%d", usersToday), float64(usersToday), "count", SeverityInfoAnalytics)
	col.add("users.newLast7Days", "New Users (7d)", "Users", fmt.Sprintf("%d", usersLast7), float64(usersLast7), "count", SeverityInfoAnalytics)
	col.add("users.newLast30Days", "New Users (30d)", "Users", fmt.Sprintf("%d", usersLast30), float64(usersLast30), "count", SeverityInfoAnalytics)

	// --- Workspaces ------------------------------------------------------------
	workspacesTotal := countRows(tx, &abacdefs.WorkspaceEntity{}, "")
	workspacesLast30 := countRows(tx, &abacdefs.WorkspaceEntity{}, "created_at >= ?", last30Days)
	workspaceTypesTotal := countRows(tx, &abacdefs.WorkspaceTypeEntity{}, "")
	membershipsTotal := countRows(tx, &abacdefs.UserWorkspaceEntity{}, "")

	col.add("workspaces.total", "Total Workspaces", "Workspaces", fmt.Sprintf("%d", workspacesTotal), float64(workspacesTotal), "count", SeverityInfoAnalytics)
	col.add("workspaces.newLast30Days", "New Workspaces (30d)", "Workspaces", fmt.Sprintf("%d", workspacesLast30), float64(workspacesLast30), "count", SeverityInfoAnalytics)
	col.add("workspaces.typesTotal", "Workspace Types", "Workspaces", fmt.Sprintf("%d", workspaceTypesTotal), float64(workspaceTypesTotal), "count", SeverityInfoAnalytics)
	col.add("memberships.total", "Workspace Memberships", "Workspaces", fmt.Sprintf("%d", membershipsTotal), float64(membershipsTotal), "count", SeverityInfoAnalytics)

	avgPerUser := safeDivide(float64(membershipsTotal), float64(usersTotal))
	avgPerWorkspace := safeDivide(float64(membershipsTotal), float64(workspacesTotal))
	col.add("memberships.avgPerUser", "Avg Workspaces / User", "Workspaces", fmt.Sprintf("%.1f", avgPerUser), avgPerUser, "count", SeverityInfoAnalytics)
	col.add("memberships.avgPerWorkspace", "Avg Members / Workspace", "Workspaces", fmt.Sprintf("%.1f", avgPerWorkspace), avgPerWorkspace, "count", SeverityInfoAnalytics)

	// --- Access & security -------------------------------------------------
	passportsTotal := countRows(tx, &abacdefs.PassportEntity{}, "")
	passportsConfirmed := countRows(tx, &abacdefs.PassportEntity{}, "confirmed = ?", true)
	passportsEmail := countRows(tx, &abacdefs.PassportEntity{}, "type = ?", PASSPORT_METHOD_EMAIL)
	passportsPhone := countRows(tx, &abacdefs.PassportEntity{}, "type = ?", PASSPORT_METHOD_PHONE)
	passportsTotp := countRows(tx, &abacdefs.PassportEntity{}, "totp_confirmed = ?", true)
	sessionsActive := countRows(tx, &abacdefs.TokenEntity{}, "valid_until > ?", now)
	rolesTotal := countRows(tx, &abacdefs.RoleEntity{}, "")
	invitesPending := countRows(tx, &abacdefs.PendingWorkspaceInviteEntity{}, "")
	joinKeysTotal := countRows(tx, &abacdefs.PublicJoinKeyEntity{}, "")

	confirmedRate := safeDivide(float64(passportsConfirmed), float64(passportsTotal)) * 100
	totpRate := safeDivide(float64(passportsTotp), float64(passportsTotal)) * 100

	col.add("passports.total", "Total Passports", "Access", fmt.Sprintf("%d", passportsTotal), float64(passportsTotal), "count", SeverityInfoAnalytics)
	col.add("passports.confirmedRate", "Passport Confirmation Rate", "Access", fmt.Sprintf("%.1f%%", confirmedRate), confirmedRate, "percent", countRateSeverity(confirmedRate))
	col.add("passports.emailCount", "Email Passports", "Access", fmt.Sprintf("%d", passportsEmail), float64(passportsEmail), "count", SeverityInfoAnalytics)
	col.add("passports.phoneCount", "Phone Passports", "Access", fmt.Sprintf("%d", passportsPhone), float64(passportsPhone), "count", SeverityInfoAnalytics)
	col.add("passports.totpAdoptionRate", "2FA Adoption Rate", "Access", fmt.Sprintf("%.1f%%", totpRate), totpRate, "percent", countRateSeverity(totpRate))
	col.add("sessions.active", "Active Sessions", "Access", fmt.Sprintf("%d", sessionsActive), float64(sessionsActive), "count", SeverityInfoAnalytics)
	col.add("roles.total", "Total Roles", "Access", fmt.Sprintf("%d", rolesTotal), float64(rolesTotal), "count", SeverityInfoAnalytics)
	col.add("invites.pending", "Pending Workspace Invites", "Access", fmt.Sprintf("%d", invitesPending), float64(invitesPending), "count", SeverityInfoAnalytics)
	col.add("joinKeys.total", "Public Join Keys", "Access", fmt.Sprintf("%d", joinKeysTotal), float64(joinKeysTotal), "count", SeverityInfoAnalytics)

	// --- Engagement ----------------------------------------------------------
	notificationsLast30 := countRows(tx, &abacdefs.NotificationEntity{}, "created_at >= ?", last30Days)
	notificationsRead := countRows(tx, &abacdefs.NotificationEntity{}, "created_at >= ? and is_read = ?", last30Days, true)
	rootAdmins := countRows(tx, &abacdefs.UserWorkspaceEntity{}, "workspace_id = ?", ROOT_VAR)

	readRate := safeDivide(float64(notificationsRead), float64(notificationsLast30)) * 100

	col.add("notifications.sentLast30Days", "Notifications Sent (30d)", "Engagement", fmt.Sprintf("%d", notificationsLast30), float64(notificationsLast30), "count", SeverityInfoAnalytics)
	col.add("notifications.readRate", "Notification Read Rate (30d)", "Engagement", fmt.Sprintf("%.1f%%", readRate), readRate, "percent", countRateSeverity(readRate))
	col.add("users.rootAdmins", "Root Workspace Admins", "Engagement", fmt.Sprintf("%d", rootAdmins), float64(rootAdmins), "count", SeverityInfoAnalytics)

	// --- Charts ----------------------------------------------------------------
	col.addSeries("users.monthlySignups", "New Users / Month", "Users", "bar", "count",
		monthlySeries(tx, &abacdefs.UserEntity{}, 12))
	col.addSeries("workspaces.monthlyCreated", "New Workspaces / Month", "Workspaces", "bar", "count",
		monthlySeries(tx, &abacdefs.WorkspaceEntity{}, 12))
	col.addSeries("workspaces.byType", "Workspaces by Type", "Workspaces", "bar", "count",
		workspacesByType(tx))
	col.addSeries("passports.byType", "Passports by Type", "Access", "bar", "count",
		[]abacdefs.AnalyticsOverviewActionResSeriesPoints{
			{Label: "email", Value: float64(passportsEmail)},
			{Label: "phone", Value: float64(passportsPhone)},
		})
	col.addSeries("roles.topByAssignment", "Most-Assigned Roles", "Access", "bar", "count",
		topRolesByAssignment(tx, 8))

	return &abacdefs.AnalyticsOverviewActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.AnalyticsOverviewActionRes{
			GeneratedAt: now.Format(time.RFC3339),
			Items:       emigo.ArrayReplace(col.items),
			Series:      emigo.ArrayReplace(col.series),
		}),
	}, nil
}

// countRows runs a plain COUNT(*) against model's table, optionally filtered - a
// bare where of "" runs no filter at all (COUNT of every row). Errors are treated as
// zero rather than failing the whole snapshot: one bad count shouldn't blank out
// every other stat tile on the page.
func countRows(tx *gorm.DB, model interface{}, where string, args ...interface{}) int64 {
	var n int64
	q := tx.Model(model)
	if where != "" {
		q = q.Where(where, args...)
	}
	q.Count(&n)
	return n
}

func safeDivide(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// monthlySeries buckets model's own createdAt column into `months` calendar months
// (oldest first, current month last) - fetched as raw timestamps and bucketed in Go
// rather than a vendor-specific date_trunc/DATE_FORMAT SQL call, so this works the
// same on every DB vendor fireback supports (see modules/fireback/dbdsn). Fine at
// admin-dashboard scale; revisit with a real GROUP BY if a table here ever grows
// large enough for that to matter.
func monthlySeries(tx *gorm.DB, model interface{}, months int) []abacdefs.AnalyticsOverviewActionResSeriesPoints {
	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, -(months - 1), 0)

	var rows []struct {
		CreatedAt time.Time
	}
	tx.Model(model).Select("created_at").Where("created_at >= ?", start).Find(&rows)

	counts := map[string]float64{}
	for _, r := range rows {
		counts[r.CreatedAt.Format("2006-01")]++
	}

	points := make([]abacdefs.AnalyticsOverviewActionResSeriesPoints, 0, months)
	for cursor := start; !cursor.After(now); cursor = cursor.AddDate(0, 1, 0) {
		key := cursor.Format("2006-01")
		points = append(points, abacdefs.AnalyticsOverviewActionResSeriesPoints{
			Label: key,
			Value: counts[key],
		})
	}
	return points
}

// workspacesByType counts WorkspaceEntity rows per WorkspaceTypeEntity, resolving
// typeId (a uniqueId, not the numeric primary key - see WorkspaceEntity.TypeId) to
// that type's own Title for display. A typeId with no matching row (a type deleted
// out from under existing workspaces) is labeled "unknown" rather than dropped, so
// the total still reconciles with workspaces.total above.
func workspacesByType(tx *gorm.DB) []abacdefs.AnalyticsOverviewActionResSeriesPoints {
	var counts []struct {
		TypeId string
		Cnt    int64
	}
	tx.Model(&abacdefs.WorkspaceEntity{}).
		Select("type_id, count(*) as cnt").
		Group("type_id").
		Order("cnt desc").
		Scan(&counts)

	if len(counts) == 0 {
		return []abacdefs.AnalyticsOverviewActionResSeriesPoints{}
	}

	ids := make([]string, 0, len(counts))
	for _, c := range counts {
		ids = append(ids, c.TypeId)
	}
	var types []abacdefs.WorkspaceTypeEntity
	tx.Where("unique_id IN ?", ids).Find(&types)
	titleById := map[string]string{}
	for _, t := range types {
		titleById[t.UniqueId] = t.Title
	}

	points := make([]abacdefs.AnalyticsOverviewActionResSeriesPoints, 0, len(counts))
	for _, c := range counts {
		label := titleById[c.TypeId]
		if label == "" {
			// Two different reasons a title can be blank here, kept visually distinct
			// so they don't collapse into one misleadingly-merged "unknown" bar: no
			// typeId at all (legacy/malformed data), vs. a real WorkspaceTypeEntity row
			// whose own Title is blank (true today for the bootstrap-created "root"
			// type, which has no user-facing title/slug by design) - falling back to
			// the raw typeId is still unique and more informative than a shared label.
			if c.TypeId == "" {
				label = "(none)"
			} else {
				label = c.TypeId
			}
		}
		points = append(points, abacdefs.AnalyticsOverviewActionResSeriesPoints{Label: label, Value: float64(c.Cnt)})
	}
	return points
}

// topRolesByAssignment counts WorkspaceRoleEntity rows (one per user-in-workspace
// role grant) per RoleEntity, the top `limit` roles by how many grants reference
// them - the roles actually doing the most work across the install, not just the
// most roles that happen to exist (see roles.total, a plain count of RoleEntity
// itself, for that).
func topRolesByAssignment(tx *gorm.DB, limit int) []abacdefs.AnalyticsOverviewActionResSeriesPoints {
	var counts []struct {
		RoleId string
		Cnt    int64
	}
	tx.Model(&abacdefs.WorkspaceRoleEntity{}).
		Select("role_id, count(*) as cnt").
		Group("role_id").
		Order("cnt desc").
		Limit(limit).
		Scan(&counts)

	if len(counts) == 0 {
		return []abacdefs.AnalyticsOverviewActionResSeriesPoints{}
	}

	ids := make([]string, 0, len(counts))
	for _, c := range counts {
		ids = append(ids, c.RoleId)
	}
	var roles []abacdefs.RoleEntity
	tx.Where("unique_id IN ?", ids).Find(&roles)
	nameById := map[string]string{}
	for _, r := range roles {
		nameById[r.UniqueId] = r.Name
	}

	points := make([]abacdefs.AnalyticsOverviewActionResSeriesPoints, 0, len(counts))
	for _, c := range counts {
		label := nameById[c.RoleId]
		if label == "" {
			label = "unknown"
		}
		points = append(points, abacdefs.AnalyticsOverviewActionResSeriesPoints{Label: label, Value: float64(c.Cnt)})
	}
	return points
}
