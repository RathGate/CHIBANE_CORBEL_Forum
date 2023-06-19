fetch('https://pokemonblog.com/feed/')
    .then(response => response.text())
    .then(data => {
        const parser = new DOMParser();
        const xmlDoc = parser.parseFromString(data, 'application/xml');

        const items = xmlDoc.getElementsByTagName('item');

        const newsElements = [];

        for (let i = 0; i < items.length; i++) {
            const item = items[i];

            const title = item.getElementsByTagName('title')[0].textContent;
            const image = item.getElementsByTagName('media:thumbnail')[0].getAttribute('url');
            const date = item.getElementsByTagName('pubDate')[0].textContent;

            const titleElement = document.createElement('h2');
            titleElement.textContent = title;
            titleElement.classList.add('news-title');

            const imageElement = document.createElement('img');
            imageElement.src = image;

            const dateElement = document.createElement('p');
            dateElement.textContent = date;
            dateElement.classList.add('news-date');

            const newsContainer = document.createElement('div');
            newsContainer.classList.add('news-item');
            newsContainer.appendChild(imageElement);
            newsContainer.appendChild(titleElement);
            newsContainer.appendChild(dateElement);

            newsElements.push(newsContainer);
        }

        const feedContainer = document.querySelector('.news-container');
        newsElements.forEach(newsElement => {
            feedContainer.appendChild(newsElement);
        });

        $('.news-container').slick({
            autoplay: true,
            autoplaySpeed: 3000,
            dots: true,
            infinite: true,
            slidesToShow: 1,
            slidesToScroll: 1,
            prevArrow: '<div class="slick-prev"><i class="fa-regular fa-circle-left"></i></div>',
            nextArrow: '<div class="slick-next"><i class="fa-regular fa-circle-right"></i></div>',

        });
    })
    .catch(error => console.log(error));