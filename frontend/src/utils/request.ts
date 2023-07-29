import axios,{ AxiosInstance } from "axios";
import notification from "ant-design-vue";
import router from "@/router";
import i18n from "@/locales";
import store from "@/stores";

interface Resp<T> {
    code:number;
    text:string;
    payload:T;
}

const { t } = i18n.global;

const ACCESS_TOKEN = sessionStorage.getItem("jwt");

const BASE_URL = "/api/v1";

const request:AxiosInstance = axios.create({
    timeout:200000,
    headers:{
        "Content-Type":"application/json",
    },
});

const responseInject = (resp:Resp<never> => {
    if (resp.text !== "" && resp.code === 1200){
        notification.info({
            message:t("common.session.state")+":1200",
            description: resp.text,
        });
    }
    if ( resp.code >1200){
        notification.error({
            message:t("common.session.state")+ `:${resp.code}`,
            description: resp.text,
        });
    }
}) ;

const overrideHeaders = () => {
    request.defaults.headers.common["Authorization"] =
        "Bearer " + store.state.user.account.token;
};

export {request, BASE_URL,overrideHeaders,Resp }