<script setup lang="ts">
import type { FormSubmitEvent } from "@nuxt/ui";
import * as z from "zod";

type PageProps = {
  todoGroup: TodoGroup;
};

const { todoGroup } = defineProps<PageProps>();

const startInput = ref(false);
const inputRef = useTemplateRef("inputRef");
const cardRef = useTemplateRef("cardRef");

onClickOutside(cardRef, () => {
  startInput.value = false;
});

async function cardFocus() {
  startInput.value = true;
  await nextTick();
  inputRef.value?.inputRef?.focus();
}

const formTodoSchema = z.object({
  task: z.string("Input task").min(1, "Input task"),
});

type FormTodoSchema = z.output<typeof formTodoSchema>;

const formTodo = reactive<Partial<FormTodoSchema>>({
  task: "",
});

async function submitTodo(e: FormSubmitEvent<FormTodoSchema>) {
  await TodoService.AddTodo(todoGroup.ID, e.data.task);
  formTodo.task = "";
}
</script>

<template>
  <div
    ref="cardRef"
    class="bg-primary-800 p-3 rounded-sm flex flex-col gap-2"
    @click="cardFocus"
  >
    <h1 class="text-primary-100 text-xl">{{ todoGroup.Title }}</h1>
    <USeparator :ui="{ border: 'border-primary-900' }" size="xs" />
    <UCheckbox v-for="todo in todoGroup.Todos" :label="todo.Task" />
    <UForm
      :schema="formTodoSchema"
      :state="formTodo"
      class="space-y-4 w-full"
      @submit="submitTodo"
    >
      <UInput
        v-if="startInput"
        ref="inputRef"
        v-model="formTodo.task"
        class="block"
        variant="none"
        :ui="{
          base: 'px-0',
        }"
      />
    </UForm>
  </div>
</template>
