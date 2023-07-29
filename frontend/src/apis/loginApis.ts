import { request } from "@/utils/request";

export function serverRedisterState() {
  return request.get("");
}
