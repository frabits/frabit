<template>
  <div class="login_page">
    <div style="margin-left: 2%">
      <a-space>
        <span>Version 2.1.19</span>
      </a-space>
    </div>
  </div>
  <div>
    <LoginForm />
  </div>
</template>

<script setup lang="ts">
import { UnwrapRef, reactive, ref, onMounted, computed } from "vue";
import router from "@/router";
import { useRoute } from "vue-router";
import LoginForm from "@/views/auth/LoginForm.vue";
import {getOIDCMetadata} from "@/apis/loginApis";


const route = useRoute();
const oidcEnabled = ref(false);
const oidcLoginUrl = ref("");
const query = computed(() => route.query).value;


</script>
