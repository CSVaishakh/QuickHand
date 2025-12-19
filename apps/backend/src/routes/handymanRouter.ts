import { eq } from "drizzle-orm";
import { db } from "../../../../packages/db";
import { user as userTable } from "../../../../packages/db/schema" ;
import { type signupRequest, type user, type Variables } from "../lib/types/types";
import { Hono } from "hono";
import { requireAuth } from "../../../../packages/auth/auth-middleware";
import { fetchJobs, findRecordsInJobs } from "../lib/queries";
import { getJobCategory } from "../lib/utils";

const handymanRouter = new Hono<{Variables: Variables}>();



handymanRouter.post('/sign-up', async (c) => {
    const body: signupRequest = await c.req.json();
    const { email, password, name} = body;
    
    const origin = new URL(c.req.url).origin;
    const response  = await fetch( `${origin}/auth/sign-up/email`, {
        method: "POST",
        headers: {
            'Content-Type' : "application/json",
        },
        body: JSON.stringify({
            email,
            password,
            name,
        })
    });

    const data = await response.json() as {token: string, user: user};
    const id  = data.user.id;
    const token = data.token;

    if(data) {
        const [updatedUser] = await db
            .update(userTable)
            .set({
                role: 'handyman',
                category: 'plumber'
            })
            .where(eq(userTable.id, id))
            .returning()

        return c.json({
            token: data.token,
            user : updatedUser
        });
    };

    return c.json({error: 'Signup Failed' }, 400);
});

handymanRouter.get('/dashboard',requireAuth, async (c) => {
    const user = c.var.user;

    const jobs = await findRecordsInJobs('handyman', user.id)

    const completeJobs = jobs.filter(job => job.job_status === "Completed");
    const incompleteJobs = jobs.filter(job => job.job_status === "NotCompleted");

    const completeJobRevenue = completeJobs.reduce((sum,job) => sum + job.cost, 0) 
    const incompleteJobRevenue = incompleteJobs.reduce((sum,job) => sum + job.cost, 0) 

    return c.json({completeJobs, completeJobRevenue, incompleteJobs, incompleteJobRevenue});
});

handymanRouter.get('/jobs-available', requireAuth, async (c) => {
    const user = c.var.user;

    const user_category = user.category;
    const job_category = getJobCategory(user_category);

    const jobs = fetchJobs(job_category)
    return c.json({jobs}, 200)
})
export default handymanRouter;