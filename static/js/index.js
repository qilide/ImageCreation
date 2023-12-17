// 展示主页动态图片
$(function () {
    $.ajax({
        url: '/image/index', // 向后端请求图片信息的路由
        method: 'GET',
        success: function(response) {
            const images = response.data; // 假设后端返回一个名为 "images" 的数组，包含图片信息
            const imageGallery = $('#imageGallery');
            console.log(images);
            // 在页面上动态生成图片展示区域
            images.forEach(imageInfo => {
                const imageUrl = imageInfo.path;
                const imageName = imageInfo.image_name;
                const imageId = imageInfo.id;
                const galleryItem = $('<div>').addClass('col-xl-3 col-lg-4 col-md-6');

                const galleryContent = `
            <div class="gallery-item h-100">
              <img src="${imageUrl}" style="max-width:100%;height:100%" alt="${imageName}">
              <div class="gallery-links d-flex align-items-center justify-content-center">
                <a href="${imageUrl}" title="${imageName}" class="glightbox preview-link"><i class="bi bi-arrows-angle-expand"></i></a>
                <a href="/gallery-single" class="details-link"><i class="bi bi-link-45deg"></i></a>
              </div>
            </div>
          `;
                galleryItem.html(galleryContent);
                imageGallery.append(galleryItem);
            });
        },
        error: function(xhr, status, error) {
            console.error('Error:', error);
        }
    });
});
