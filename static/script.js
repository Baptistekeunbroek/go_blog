// Fetch and display all posts
function fetchPosts() {
  fetch("/posts")
    .then((response) => response.json())
    .then((posts) => {
      const postsContainer = document.getElementById("posts");
      postsContainer.innerHTML = ""; // Clear existing posts

      for (const id in posts) {
        const post = posts[id];
        console.log(post, "post");
        postsContainer.innerHTML += `
                    <div class="post-card">
                        <div class="post-content">
                            <h3 class="post-title">${post.Title}</h3>
                            <p class="post-author"><strong>Author:</strong> ${
                              post.Author
                            }</p>
                            <p class="post-body">${post.Content}</p>
                            <p class="post-date"><small>Created: ${new Date(
                              post.Created
                            ).toLocaleString()}</small></p>
                        </div>
                    </div>
                `;
      }
    });
}

// Create a new post
document
  .getElementById("create-post-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();

    const title = document.getElementById("title").value;
    const author = document.getElementById("author").value;
    const content = document.getElementById("content").value;

    fetch("/posts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, author, content }),
    })
      .then((response) => response.json())
      .then(() => {
        fetchPosts(); // Refresh the posts list
        document.getElementById("create-post-form").reset(); // Clear form
      });
  });

// Update an existing post
document
  .getElementById("update-post-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();

    const id = document.getElementById("update-id").value;
    const title = document.getElementById("update-title").value;
    const author = document.getElementById("update-author").value;
    const content = document.getElementById("update-content").value;

    fetch(`/posts/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, author, content }),
    })
      .then((response) => response.json())
      .then(() => {
        fetchPosts(); // Refresh the posts list
        document.getElementById("update-post-form").reset(); // Clear form
      });
  });

// Delete a post
document
  .getElementById("delete-post-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();

    const id = document.getElementById("delete-id").value;

    fetch(`/posts/${id}`, {
      method: "DELETE",
    }).then(() => {
      fetchPosts(); // Refresh the posts list
      document.getElementById("delete-post-form").reset(); // Clear form
    });
  });

// Fetch and display posts when the page loads
window.onload = fetchPosts;
