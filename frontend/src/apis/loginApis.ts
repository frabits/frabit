import { request } from "@/utils/request";

export function serverRegisterState() {
  return request.get("");
}

export function getOIDCMetadata() {
  return request.get("");
}

export interface LoginForm {
  username: "",
  password: "",
  isLDAP: false,
  isOIDC: false,
}