<script setup lang="ts">
import TodoTask from "./TodoTask.vue";

type PageProps = {
  todoGroup: TodoGroup;
};

const { todoGroup } = defineProps<PageProps>();

const taskRefs = ref(new Map());
const startInput = ref(false);
const cardRef = useTemplateRef("cardRef");

function setTaskRef(el: any, id: number) {
  if (el) taskRefs.value.set(id, el);
  else taskRefs.value.delete(id);
}

const onTaskCreated = async (newTodo: Todo) => {
  todoGroup.Todos?.push(newTodo);
  await nextTick();
  taskRefs.value.get(newTodo.ID)?.focus();
};

async function deleteTodoGroup() {
  await TodoService.DeleteTodoGroup(todoGroup.ID);
}

onClickOutside(cardRef, () => {
  startInput.value = false;
});
</script>

<template>
  <div ref="cardRef" class="bg-primary-800 p-3 rounded-sm flex flex-col gap-3">
    <div class="w-full flex items-center justify-between gap-1">
      <h1 class="text-primary-100 text-xl">{{ todoGroup.Title }}</h1>
      <UButton
        icon="i-ic-baseline-delete"
        variant="soft"
        @click="deleteTodoGroup"
      />
    </div>
    <USeparator :ui="{ border: 'border-primary-900' }" size="xs" />
    <div>
      <TodoTask
        :ref="(el) => setTaskRef(el, todo.ID)"
        :key="todo.ID"
        :todo="todo"
        v-for="todo in todoGroup.Todos"
        edit-mode
      />
      <TodoTask
        :todo="{
          ID: 0,
          Finished: false,
          Task: '',
          TodoGroupID: todoGroup.ID,
        }"
        placeholder="New task"
        @created="onTaskCreated"
      />
    </div>
  </div>
</template>
