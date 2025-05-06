<template>
  <div class="p-4 max-w-3xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Welcome Home</h1>

    <!-- Create Post -->
    <div class="mb-6">
      <h2 class="text-xl mb-2">Create Post</h2>
      <textarea v-model="newPost.content" class="w-full border p-2 rounded mb-2" rows="3" placeholder="What's on your mind?"></textarea>
      <input type="file" @change="handleImageUpload" class="mb-2" />
      <button @click="createPost" class="bg-blue-600 text-white px-4 py-2 rounded">Post</button>
    </div>

    <!-- Posts List -->
    <PostList :posts="posts" @edit="editPost" @delete="deletePost" />

    <!-- Edit Post Modal (Optional) -->
    <div v-if="isEditing" class="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div class="bg-white p-6 rounded shadow-lg">
        <h2 class="text-xl mb-4">Edit Post</h2>
        <textarea v-model="newPost.content" class="w-full border p-2 rounded mb-2" rows="3" placeholder="Update your post"></textarea>
        <input type="file" @change="handleImageUpload" class="mb-2" />
        <button @click="saveEdit" class="bg-blue-600 text-white px-4 py-2 rounded">Save</button>
        <button @click="cancelEdit" class="bg-gray-400 text-white px-4 py-2 rounded ml-2">Cancel</button>
      </div>
    </div>

    <!-- Logout -->
    <button class="mt-8 bg-red-500 text-white px-4 py-2 rounded" @click="logout">Logout</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import PostList from '../components/PostList.vue'

const newPost = ref({ content: '', image: null })
const posts = ref([])
const isEditing = ref(false)
const currentEditPost = ref(null)

const fetchPosts = async () => {
  const res = await axios.get('http://localhost:8000/posts')
  posts.value = res.data
}

const createPost = async () => {
  const formData = new FormData()
  formData.append('content', newPost.value.content)
  if (newPost.value.image) {
    formData.append('image', newPost.value.image)
  }
  await axios.post('http://localhost:8000/posts', formData)
  newPost.value = { content: '', image: null }
  fetchPosts()
}

const handleImageUpload = (e) => {
  newPost.value.image = e.target.files[0]
}

const deletePost = async (id) => {
  await axios.delete(`http://localhost:8000/posts/${id}`)
  fetchPosts()
}

const editPost = (post) => {
  currentEditPost.value = post
  newPost.value = { content: post.content, image: null }
  isEditing.value = true
}

const saveEdit = async () => {
  const formData = new FormData()
  formData.append('content', newPost.value.content)
  if (newPost.value.image) {
    formData.append('image', newPost.value.image)
  }
  
  await axios.put(`http://localhost:8000/posts/${currentEditPost.value.id}`, formData)
  isEditing.value = false
  fetchPosts()
}

const cancelEdit = () => {
  isEditing.value = false
  currentEditPost.value = null
}

const logout = () => {
  localStorage.clear()
  window.location.href = '/login'
}

onMounted(fetchPosts)
</script>
