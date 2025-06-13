<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';

interface Todo {
  id: string;
  title: string;
  completed: boolean;
}

interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

const todos = ref<Todo[]>([]);
const users = ref<User[]>([]);
const selectedUser = ref<User | null>(null);
const showUserForm = ref(false);
const userForm = ref({
  name: '',
  email: '',
  editing: false,
  editingId: null as number | null
});

onBeforeMount(async () => {
  await loadTodos();
  await loadUsers();
});

const loadTodos = async () => {
  try {
    const response = await fetch('http://localhost:8080/todos');
    const json = await response.json();
    todos.value = json.todos;
  } catch (error) {
    console.error('Error loading todos:', error);
  }
};

const loadUsers = async () => {
  try {
    const response = await fetch('http://localhost:8080/users');
    const json = await response.json();
    users.value = json.users;
  } catch (error) {
    console.error('Error loading users:', error);
  }
};

const showUser = async (userId: number) => {
  try {
    const response = await fetch(`http://localhost:8080/users/${userId}`);
    const json = await response.json();
    selectedUser.value = json.user;
  } catch (error) {
    console.error('Error loading user:', error);
  }
};

const openUserForm = (user?: User) => {
  if (user) {
    userForm.value = {
      name: user.name,
      email: user.email,
      editing: true,
      editingId: user.id
    };
  } else {
    userForm.value = {
      name: '',
      email: '',
      editing: false,
      editingId: null
    };
  }
  showUserForm.value = true;
};

const closeUserForm = () => {
  showUserForm.value = false;
  userForm.value = {
    name: '',
    email: '',
    editing: false,
    editingId: null
  };
};

const submitUserForm = async () => {
  try {
    const method = userForm.value.editing ? 'PUT' : 'POST';
    const url = userForm.value.editing 
      ? `http://localhost:8080/users/${userForm.value.editingId}`
      : 'http://localhost:8080/users';
    
    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: userForm.value.name,
        email: userForm.value.email
      })
    });

    if (response.ok) {
      await loadUsers();
      closeUserForm();
    } else {
      console.error('Error saving user');
    }
  } catch (error) {
    console.error('Error saving user:', error);
  }
};

</script>

<template>
  <div style="padding: 20px; max-width: 1200px; margin: 0 auto;">
    <h1>Togo Application</h1>
    
    <!-- Todos Section -->
    <div style="margin-bottom: 40px;">
      <h2>Todos</h2>
      <ul>
        <li v-for="todo in todos" :key="todo.id">
          {{ todo.title }}
        </li>
      </ul>
    </div>

    <!-- Users Section -->
    <div>
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h2>Users</h2>
        <button @click="openUserForm()" style="padding: 8px 16px; background-color: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer;">
          Add User
        </button>
      </div>
      
      <!-- Users List -->
      <div style="margin-bottom: 20px;">
        <h3>Users List</h3>
        <div v-if="users.length === 0" style="color: #666;">No users found</div>
        <div v-else style="display: grid; gap: 10px;">
          <div 
            v-for="user in users" 
            :key="user.id"
            style="border: 1px solid #ddd; padding: 15px; border-radius: 5px; background-color: #f9f9f9;"
          >
            <div style="display: flex; justify-content: space-between; align-items: center;">
              <div>
                <div style="font-weight: bold;">{{ user.name }}</div>
                <div style="color: #666;">{{ user.email }}</div>
                <div style="font-size: 0.8em; color: #999;">ID: {{ user.id }}</div>
              </div>
              <div style="display: flex; gap: 10px;">
                <button @click="showUser(user.id)" style="padding: 5px 10px; background-color: #28a745; color: white; border: none; border-radius: 3px; cursor: pointer;">
                  View
                </button>
                <button @click="openUserForm(user)" style="padding: 5px 10px; background-color: #ffc107; color: black; border: none; border-radius: 3px; cursor: pointer;">
                  Edit
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Selected User Details -->
      <div v-if="selectedUser" style="margin-bottom: 20px;">
        <h3>User Details</h3>
        <div style="border: 1px solid #ddd; padding: 20px; border-radius: 5px; background-color: #f0f8ff;">
          <div><strong>ID:</strong> {{ selectedUser.id }}</div>
          <div><strong>Name:</strong> {{ selectedUser.name }}</div>
          <div><strong>Email:</strong> {{ selectedUser.email }}</div>
          <div><strong>Created:</strong> {{ new Date(selectedUser.created_at).toLocaleString() }}</div>
          <div><strong>Updated:</strong> {{ new Date(selectedUser.updated_at).toLocaleString() }}</div>
          <button @click="selectedUser = null" style="margin-top: 10px; padding: 5px 10px; background-color: #6c757d; color: white; border: none; border-radius: 3px; cursor: pointer;">
            Close
          </button>
        </div>
      </div>

      <!-- User Form Modal -->
      <div v-if="showUserForm" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background-color: rgba(0,0,0,0.5); display: flex; justify-content: center; align-items: center; z-index: 1000;">
        <div style="background-color: white; padding: 30px; border-radius: 10px; width: 90%; max-width: 400px;">
          <h3>{{ userForm.editing ? 'Edit User' : 'Add New User' }}</h3>
          <form @submit.prevent="submitUserForm">
            <div style="margin-bottom: 15px;">
              <label style="display: block; margin-bottom: 5px; font-weight: bold;">Name:</label>
              <input 
                v-model="userForm.name" 
                type="text" 
                required
                style="width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; box-sizing: border-box;"
              />
            </div>
            <div style="margin-bottom: 20px;">
              <label style="display: block; margin-bottom: 5px; font-weight: bold;">Email:</label>
              <input 
                v-model="userForm.email" 
                type="email" 
                required
                style="width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; box-sizing: border-box;"
              />
            </div>
            <div style="display: flex; gap: 10px; justify-content: flex-end;">
              <button type="button" @click="closeUserForm" style="padding: 8px 16px; background-color: #6c757d; color: white; border: none; border-radius: 4px; cursor: pointer;">
                Cancel
              </button>
              <button type="submit" style="padding: 8px 16px; background-color: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer;">
                {{ userForm.editing ? 'Update' : 'Create' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
