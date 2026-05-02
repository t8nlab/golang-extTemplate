    import { log, defineAction } from "@titanpl/native";
import Extension from '@ext/go-extension';


export default defineAction((req) => {
const ext = new Extension();
const result = ext.addNumber(10, 20);
log(result); // 30
return result
})
