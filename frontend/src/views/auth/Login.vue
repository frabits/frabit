<template>
  <div class="login_page">
    <div style="margin-left: 2%">
      <a-space>
        <a-button
            v-if="enableRegister"
            type="dashed"
            ghost
            @click="isOpen = true"
        >{{$t('user.form.title') }}
        </a-button>
      </a-space>
    </div>
    <div class="login_container">
      <LoginForm />
    </div>
    <div class="footer">
      <div style="text-align: center">{{Copyright}}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, onMounted} from "vue";
import LoginForm from "@/views/auth/LoginForm.vue";

const {isOpen} = CommonMixin();
const enableRegister = ref(false)
onMounted( async () => {
  const {data} = await serverRedisterState();
  enableRegister.value = data.payload.reg
});
</script>
