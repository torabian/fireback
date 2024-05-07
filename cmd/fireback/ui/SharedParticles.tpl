{{ define "products-inline" }}
    {{ range . }}
    <div class="swiper-slide">
        <div class="product-card position-relative">
            <div class="image-holder">
            {{ if index .Image 0 }}
                {{ $img := index .Image 0}}
                <img src="http://localhost:4502/files-inline/{{ $img.DiskPath }}" alt="product-item" class="img-fluid">
            {{ end }}
            </div>
            <div class="cart-concern position-absolute">
            <div class="cart-button d-flex">
                <a href="#" class="btn btn-medium btn-black">Add to Cart<svg class="cart-outline"><use xlink:href="#cart-outline"></use></svg></a>
            </div>
            </div>
            <div class="card-detail d-flex justify-content-between align-items-baseline pt-3">
            <h3 class="card-title text-uppercase">
                <a href="#">{{ .Name }}</a>
            </h3>
            <span class="item-price text-primary">$980</span>
            </div>
        </div>
    </div>
    {{ end }}

{{ end }}