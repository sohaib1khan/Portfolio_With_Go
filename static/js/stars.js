// Wait until the DOM (Document Object Model) is fully loaded
document.addEventListener('DOMContentLoaded', () => {

    // Function to create a single shooting star
    function createShootingStar() {
        // Create a new <div> element to represent the shooting star
        const shootingStar = document.createElement('div');
        shootingStar.classList.add('shooting-star'); // Add the 'shooting-star' CSS class for styling
        document.body.appendChild(shootingStar); // Append the shooting star to the <body> of the document

        // Randomize the starting position of the shooting star
        const startX = Math.random() * window.innerWidth; // Random horizontal position across the screen width
        const startY = Math.random() * window.innerHeight / 2; // Random vertical position in the upper half of the screen

        // Randomize the ending position to create variation in the trajectory
        const endX = startX + (Math.random() * 400 - 200); // Add some horizontal variation to the movement
        const endY = startY + Math.random() * window.innerHeight; // Move downward to a random vertical position

        // Set the initial position of the shooting star using inline styles
        shootingStar.style.left = `${startX}px`; // Set the starting horizontal position
        shootingStar.style.top = `${startY}px`; // Set the starting vertical position

        // Apply the CSS animation with a random duration between 1.5 and 3.5 seconds
        shootingStar.style.animation = `shootingStarPath ${Math.random() * 2 + 1.5}s linear`;

        // Remove the shooting star from the DOM after the animation ends to prevent memory leaks
        shootingStar.addEventListener('animationend', () => {
            shootingStar.remove(); // Clean up the element
        });
    }

    // Create multiple shooting stars at regular intervals
    setInterval(() => {
        // Randomly decide how many shooting stars to create (between 10 and 15)
        const starCount = Math.floor(Math.random() * 6) + 10; // Random integer in the range [10, 15]

        // Loop to create the determined number of shooting stars
        for (let i = 0; i < starCount; i++) {
            createShootingStar(); // Create a single shooting star
        }
    }, 2000); // Repeat every 2 seconds to create a batch of shooting stars
});
