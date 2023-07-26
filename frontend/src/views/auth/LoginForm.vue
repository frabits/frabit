<template>
  <img src="../../assets/logo.png" width="359"/>
  <br />
  <a-form :model="loginForm">
    <a-form-item>
      <a-input
          v-model:value="loginForm.username"
          :placeholder="$t('user.username')"
          style="border-radius: 10px"
      />
    </a-form-item>
    <a-form-item>
      <a-input
          v-model:value="loginForm.password"
          :placeholder="$t('user.password')"
          style="border-radius: 10px"
          type="password"
          @pree-enter="() => userLogin()"
      />
    </a-form-item>
    <a-form-item>
      <a-space :size="40">
        <a-checkbox v-model:checked(v-model)="loginForm.isLDAP">
          <span class="fff">LDAP</span>
        </a-checkbox>
      </a-space>
    </a-form-item>
    <a-button type="dashed" block ghost @click="userLogin">{{$t('common.login')}}</a-button>
    <a-button
        v-if="oidcEnabled"
              type="dashed"
              block
              ghost
              style="margin-top: 10px"
              @click="oidcLogin"
    >OIDC {{$t('common.login')}}
    </a-button>
  </a-form>
</template>

<script setup lang="ts">
  import { UnwrapRef,reactive,ref,onMounted,computed } from "vue";
  import router from "@/router";
  import {useRoute} from "vue-router";
  import {debounce} from "lodash-es";

  const loginForm:UnwrapRef<LoginForm> = reactive({
    username:'',
    password:'',
    isLDAP:false,
    isOIDC:false,
      });

  const route = useRoute();
  const oidcEnabled = ref(false);
  const oidcLoginUrl = ref("");
  const query = computed(() => route.query).value;

  const getOIDC = async () =>{
    const { data } = await getOIDCMetadata();
    if ( data.code == 1200 && data.payload.enabled && data.payload.authURL) {
      oidcEnabled.value = true;
      oidcLoginUrl.value = data.payload.authURL;
    }
  };


  const oidcLogin = () => {
    window.location.href = oidcLoginUrl.value;
  };

  onMounted(() => {
    getOIDC();
    if (query.oidcLogin) {
      const {oidcLogin,...rest} = query;
      router.replace("/home");
    }
  });

  const userLogin = debounce(async () => {
    const { data} = await Login(loginForm);
    if (data.code===13001){
      return
    }
  });
</script>

