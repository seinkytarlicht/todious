<script lang="ts" setup>
import type { FormSubmitEvent } from "@nuxt/ui";
import * as z from "zod";

type PageProps = {
  todo: Todo;
  editMode?: boolean;
  placeholder?: string;
};

const { todo, editMode = false, placeholder } = defineProps<PageProps>();
const emit = defineEmits(["created"]);

const formRef = useTemplateRef("formRef");
const textareaRef = useTemplateRef("textareaRef");

const triggerSubmit = () => {
  formRef.value?.submit();
};
const debouncedTriggerSubmit = useDebounceFn(triggerSubmit, 500);

const formTodoSchema = z.object({
  groupId: z.number(),
  id: z.number(),
  task: z.string("Input task").min(1, "Input task"),
  finished: z.boolean(),
});

type FormTodoSchema = z.output<typeof formTodoSchema>;

const formTodo = reactive<Partial<FormTodoSchema>>({
  groupId: todo?.TodoGroupID ?? 0,
  id: todo?.ID ?? 0,
  task: todo?.Task ?? "",
  finished: todo?.Finished ?? false,
});

async function submitTodo(e: FormSubmitEvent<FormTodoSchema>) {
  if (editMode) {
    TodoService.UpdateTodo(e.data.id, e.data.task, e.data.finished);
  } else {
    const newTodo = await TodoService.AddTodo(e.data.groupId, e.data.task);
    formTodo.task = "";
    emit("created", newTodo);
  }
}

async function deleteTodo() {
  if (editMode) {
    await TodoService.RemoveTodo(todo.ID);
  }
}

defineExpose({
  focus: () => {
    textareaRef.value?.textareaRef?.focus?.() ??
      textareaRef.value?.textareaRef?.querySelector("textarea")?.focus();
  },
});
</script>

<template>
  <UForm
    ref="formRef"
    :schema="formTodoSchema"
    :state="formTodo"
    class="w-full"
    :disabled="false"
    @submit="submitTodo"
  >
    <div class="w-full flex items-center gap-3 border-b border-default">
      <UCheckbox
        color="neutral"
        v-model="formTodo.finished"
        @update:model-value="debouncedTriggerSubmit"
        :disabled="!editMode"
      />
      <UTextarea
        ref="textareaRef"
        :placeholder="placeholder"
        class="w-full"
        variant="none"
        :ui="{
          base: 'px-0 rounded-none',
        }"
        v-model="formTodo.task"
        @update:model-value="debouncedTriggerSubmit"
        :rows="1"
        autoresize
      />
      <UButton
        v-if="editMode"
        icon="i-ic-baseline-close"
        variant="link"
        class="p-0"
        @click="deleteTodo"
      />
    </div>
  </UForm>
</template>
