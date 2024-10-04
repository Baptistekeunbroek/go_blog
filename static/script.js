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

      // Get the username from local storage
      const username = localStorage.getItem("username"); // Get the logged-in user's username

      // Convert posts object to an array and sort by created date descending
      const sortedPosts = Object.values(posts).sort(
        (a, b) => new Date(b.created) - new Date(a.created)
      );

      // Loop through sorted posts and display them
      sortedPosts.forEach((post) => {
        postsContainer.innerHTML += `
                  <div class="post-card">
                      <div class="post-content">
                      <h3 class="post-title"> ${post.author}</h3>
                          <h4 class="post-author">${post.title}</h4>
                          
                          <img src="${
                            post.image
                          }" alt="Post Image" class="post-image">
                          <p class="post-body">${post.content}</p>
                          <p class="post-date"><small>Created: ${new Date(
                            post.created
                          ).toLocaleString()}</small></p>
                      </div>
                      <div class="post-actions">
                          <button onclick="editPost('${post.id}')">Edit</button>
                          <button onclick="deletePost('${
                            post.id
                          }')">Delete</button>
                      </div>
                  </div>
              `;
      });
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

    // Get the username from local storage in user

    if (!localStorage.getItem("user")) {
      alert("You need to log in to create a post");
      return;
    } else {
      username = JSON.parse(localStorage.getItem("user")).username;
    }

    const formData = new FormData();
    formData.append("title", document.getElementById("title").value);
    formData.append("author", username); // Directly use the username from localStorage
    formData.append("content", document.getElementById("content").value);
    formData.append("image", document.getElementById("image").files[0]);

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
  const title = prompt("New Title:");
  const content = prompt("New Content:");
  const image = document.createElement("input");
  image.type = "file";
  image.accept = "image/*";

  image.onchange = function () {
    const formData = new FormData();
    if (title) formData.append("title", title);
    if (content) formData.append("content", content);

    if (image.files.length > 0) {
      formData.append("image", image.files[0]);
    }

    fetch(`/posts/${id}`, {
      method: "PUT",
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
      })
      .catch((error) => {
        console.error("Error updating post:", error);
        alert("Failed to update post. Check the console for details.");
      });
  };

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
