import { Hono } from "hono";
import type { Variables } from "../lib/types/types";
import { SigninSchema, UpdateImageSchema } from "../lib/schemas/common.schema";
import { z } from "zod";
import { raw } from "hono/html";
import { updateImage } from "../lib/queries";

const commonRouter = new Hono<{Variables: Variables}>();

commonRouter.post('/update-image', async (c) => {
    const raw = await c.req.json();

    const result = UpdateImageSchema.safeParse(raw)
    if(!result.success){
        return c.json({ error: z.treeifyError(result.error) }, 400)
    }
    const { img, userId } = result.data;

    try{
        const [imgUpdate] = await updateImage(userId, img);
        return c.json({"message":"Update Sucessfull!"});
    }catch(error){
        return c.json(
            { error: "Update Failed" }, 500
        )
    }
});

export default commonRouter;