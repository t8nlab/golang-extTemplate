import { log } from "@titanpl/native";
import { addNumber } from "@ext/template";

export function getuser(req) {
    log("Handling user request...");
    log(addNumber(9, 8));
    return {
        message: "Hello from JavaScript action!",
        user_id: req.params.id
    };
}
