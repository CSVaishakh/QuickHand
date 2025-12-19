import { db, user as userTable, eq } from "@repo/db";
import { type listed_job, type signupRequest, type user, type Variables } from "../lib/types/types";
import { Hono } from "hono";
import { requireAuth } from "@repo/auth";
import { findRecordsInJobs, listJob } from "../lib/queries";

const customerRouter = new Hono<{Variables: Variables}>();

customerRouter.post('/sign-up', async (c) => {
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
            name
        })
    });

    const data = await response.json() as {token: string, user: user};
    const id  = data.user.id;
    const token = data.token;

    if(data) {
        const [updatedUser] = await db
            .update(userTable)
            .set({ role: 'customer'})
            .where(eq(userTable.id, id))
            .returning()

        return c.json({
            token: data.token,
            user : updatedUser
        })
    }

    return c.json({ error: 'Signup Failed' }, 400 );
})

customerRouter.get('/dashboard',requireAuth, async (c) => {
    const user = c.var.user

    const jobs = await findRecordsInJobs('customer', user.id)

    const completeJobs = jobs.filter(job => job.job_status === "Completed");
    const incompleteJobs = jobs.filter(job => job.job_status === "NotCompleted");

    const completeJobRevenue = completeJobs.reduce((sum,job) => sum + job.cost, 0) 
    const incompleteJobRevenue = incompleteJobs.reduce((sum,job) => sum + job.cost, 0)
    
    const handymen = jobs.map(job => job.handyman)

    c.set("jobs",jobs)

    return c.json({completeJobs, completeJobRevenue, incompleteJobs, incompleteJobRevenue, handymen});
})

customerRouter.get('/create-job', (c) => {
    

    return c.json({})
})

customerRouter.post('/list-job', async (c) => {

    const body: listed_job = await c.req.json();
    const { id, name, customer, pay_range, job_category} = body;

    const job: listed_job = {
        id: id,
        name: name,
        customer: customer,
        pay_range: pay_range,
        job_category: job_category
    };

    const response = await listJob(job);

    if (!response){
        return c.json({error: "Job not listed"}, 400);
    }

    return c.json({
        message: "Job listed successfully",
        response
    }, 200);
})


export default customerRouter;