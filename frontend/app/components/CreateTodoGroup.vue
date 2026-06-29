<script setup lang="ts">
import * as z from "zod";
import type { FormSubmitEvent } from "@nuxt/ui";

const input = useTemplateRef("input");

const formTodoGroupSchema = z.object({
  title: z.string("Input Title").min(1, "Input title"),
});

type FormTodoGroupSchema = z.output<typeof formTodoGroupSchema>;

const formTodoGroup = reactive<Partial<FormTodoGroupSchema>>({
  title: "",
});

async function submitTodoGroup(e: FormSubmitEvent<FormTodoGroupSchema>) {
  await TodoService.CreateTodoGroup(e.data.title);
  formTodoGroup.title = "";
}

defineShortcuts({
  "/": () => {
    input.value?.inputRef?.focus();
  },
});
</script>

<template>
  <UForm
    :schema="formTodoGroupSchema"
    :state="formTodoGroup"
    class="space-y-4 w-full"
    @submit="submitTodoGroup"
  >
    <UInput
      ref="input"
      name="title"
      placeholder="Create new group task"
      class="w-full"
      size="xl"
      :ui="{
        base: 'rounded-full py-2 px-4',
        trailing: 'pr-1 gap-2',
      }"
      v-model="formTodoGroup.title"
    >
      <template #trailing>
        <UKbd value="/" />
        <UButton
          icon="i-ic-baseline-arrow-forward"
          class="rounded-full"
          type="submit"
        />
      </template>
    </UInput>
  </UForm>
</template>
