import { db, user as userTable, eq } from "@repo/db";
import { type accepct_job, type handymenSignupRequest, type job, type listed_job, type user, type Variables } from "../lib/types/types";
import { Hono } from "hono";
import { requireAuth } from "@repo/auth";
import { accepctJob, fetchJobs, findRecordsInJobs } from "../lib/queries";
import { getJobCategory } from "../lib/utils";

const handymanRouter = new Hono<{Variables: Variables}>();

handymanRouter.post('/sign-up', async (c) => {
    const body: handymenSignupRequest = await c.req.json();
    const { email, password, name, category} = body;
    
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
                category: category
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

handymanRouter.post('/accepct-job', requireAuth, async (c) => {
    const user = c.var.user;
    const req_data = await c.req.json();
    const { list_id, job_name, customer, pay_range, job_category } = req_data;

    const job: accepct_job = {
        id: crypto.randomUUID(),
        name: job_name,
        customer: customer,
        handyman: user.id,
        cost: Number(pay_range),
        job_status: "NotCompleted",
        list_id: list_id,
        job_category: job_category
    };


    const response = await accepctJob(job);

    if(!response){
        return c.json({error: "Job not accepcted"}, 400);
    }
    
    return c.json({
        message: "Job Accepted",
        response
    }, 200);
})

export default handymanRouter;