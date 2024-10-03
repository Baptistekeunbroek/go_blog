// Fetch and display all posts
function fetchPosts() {
  fetch("/posts")
    .then((response) => {
      if (!response.ok) {
        return response.text().then((text) => {
          throw new Error(`Error: ${text}`);
        });
      }
      return response.json();
    })
    .then((posts) => {
      const postsContainer = document.getElementById("posts");
      postsContainer.innerHTML = ""; // Clear existing posts

      for (const id in posts) {
        const post = posts[id];
        postsContainer.innerHTML += `
          <div class="post-card">
            <div class="post-content">
              <h3 class="post-title"><strong>${post.title}</strong></h3>
              <p class="post-author">${post.author}</p>
              <img src="${post.image}" alt="Post Image" class="post-image">
              <p class="post-body">${post.content}</p>
              <p class="post-date"><small>Created: ${new Date(
                post.created
              ).toLocaleString()}</small></p>
            </div>
            <div class="post-actions">
              <button onclick="editPost('${id}')">Edit</button>
              <button onclick="deletePost('${id}')">Delete</button>
            </div>
          </div>
        `;
      }
    })
    .catch((error) => {
      console.error("Error fetching posts:", error);
      alert("Failed to fetch posts. Check the console for details.");
    });
}

// Create a new post
document
  .getElementById("create-post-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();

    const formData = new FormData();
    formData.append("title", document.getElementById("title").value);
    formData.append("author", document.getElementById("author").value);
    formData.append("content", document.getElementById("content").value);
    formData.append("image", document.getElementById("image").files[0]); // Append image

    fetch("/posts", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (!response.ok) {
          return response.text().then((text) => {
            throw new Error(`Error: ${text}`);
          });
        }
        return response.json();
      })
      .then(() => {
        fetchPosts(); // Refresh the posts list
        document.getElementById("create-post-form").reset(); // Clear form
      })
      .catch((error) => {
        console.error("Error creating post:", error);
        alert("Failed to create post. Check the console for details.");
      });
  });

// Update an existing post
function editPost(id) {
  // For simplicity, we can prompt the user for new data (replace with modal for better UX)
  const title = prompt("New Title:");
  const author = prompt("New Author:");
  const content = prompt("New Content:");
  const image = document.createElement("input");
  image.type = "file";
  image.accept = "image/*";

  // Display an image file selector
  image.onchange = function () {
    const formData = new FormData();
    if (title) formData.append("title", title);
    if (author) formData.append("author", author);
    if (content) formData.append("content", content);

    // If an image is selected, add it to the form data
    if (image.files.length > 0) {
      formData.append("image", image.files[0]);
    }

    // Send the updated post data (text and image)
    fetch(`/posts/${id}`, {
      method: "PUT",
      body: formData, // Use FormData to handle text and file uploads
    })
      .then((response) => {
        if (!response.ok) {
          return response.text().then((text) => {
            throw new Error(`Error: ${text}`);
          });
        }
        return response.json();
      })
      .then(() => {
        fetchPosts(); // Refresh the posts list
      })
      .catch((error) => {
        console.error("Error updating post:", error);
        alert("Failed to update post. Check the console for details.");
      });
  };

  // Trigger the file selector
  image.click();
}

// Delete a post
function deletePost(id) {
  if (confirm("Are you sure you want to delete this post?")) {
    fetch(`/posts/${id}`, {
      method: "DELETE",
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to delete the post");
        }
        fetchPosts(); // Refresh the posts list
      })
      .catch((error) => {
        console.error("Error deleting post:", error);
        alert("Failed to delete post. Check the console for details.");
      });
  }
}

// Initial fetch of posts
fetchPosts();
