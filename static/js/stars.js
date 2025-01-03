document.addEventListener('DOMContentLoaded', () => {
    function createShootingStar() {
        const shootingStar = document.createElement('div');
        shootingStar.classList.add('shooting-star');
        document.body.appendChild(shootingStar);

        // Randomize the start position and duration
        const startX = Math.random() * window.innerWidth;
        const startY = Math.random() * window.innerHeight / 2;
        const endX = startX + (Math.random() * 400 - 200); // Slight horizontal variation
        const endY = startY + Math.random() * window.innerHeight;

        shootingStar.style.left = `${startX}px`;
        shootingStar.style.top = `${startY}px`;

        shootingStar.style.animation = `shootingStarPath ${Math.random() * 2 + 1.5}s linear`;

        // Remove the shooting star after animation ends
        shootingStar.addEventListener('animationend', () => {
            shootingStar.remove();
        });
    }

    // Create 10–15 shooting stars every 2 seconds
    setInterval(() => {
        const starCount = Math.floor(Math.random() * 6) + 10; // Random number between 10 and 15
        for (let i = 0; i < starCount; i++) {
            createShootingStar();
        }
    }, 2000); // Adjust the interval to make stars appear more frequently
});
