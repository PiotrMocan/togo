<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';
const todos = ref<{ id: string, title: string, completed: boolean }[]>([]);
onBeforeMount(async () => {
  const response = await fetch('http://localhost:8080/todos');
  const json = await response.json();
  todos.value = json.todos;
});

</script>

<template>
  <div>
    <h1>Todos</h1>
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        {{ todo.title }}
      </li>
    </ul>
  </div>
</template>

<style scoped></style>
