document.addEventListener("DOMContentLoaded", () => {
    fetch('/api/resources')
        .then(response => response.json())
        .then(data => {
            const container = document.getElementById('resource-container');
            data.forEach(resource => {
                const resourceDiv = document.createElement('div');
                resourceDiv.className = 'resource';
                resourceDiv.innerHTML = `
                    <img src="${resource.icon}" alt="${resource.title} Icon">
                    <h3>${resource.title}</h3>
                    <p>${resource.description}</p>
                    <a href="${resource.url}" target="_blank">Visit</a>
                `;
                container.appendChild(resourceDiv);
            });
        });
});
