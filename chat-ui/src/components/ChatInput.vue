<template>
  <form class="input-section" v-on:submit.prevent="addNewMessage">
    <InputText v-model="message" class="input-message" />
    <Button
      type="submit"
      label="Enviar"
      class="send-button p-button-outlined"
      :icon="icon"
    />
  </form>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from "vue";

export default defineComponent({
  name: "chat-input",
  emits: ["newMessage"],
  props: {
    loading: Boolean,
    error: String,
  },
  setup(props, context) {
    const message = ref("");
    const icon = ref("pi pi-send");

    const addNewMessage = () => {
      context.emit("newMessage", message.value);
    };

    // Show loading icon when the data is being send.
    watch(
      () => props.loading,
      () => {
        if (props.loading) {
          icon.value = "pi pi-spin pi-spinner";
        } else {
          icon.value = "pi pi-send";
          if (message.value && !props.error) message.value = "";
        }
      }
    );

    return { message, addNewMessage, icon };
  },
});
</script>

<style lang="scss" scoped>
.input-section {
  width: 100%;
  display: flex;
  padding: 0.5em;

  .input-message {
    width: 100%;
  }
  .send-button {
    margin-left: 1em;
  }
}
</style>
