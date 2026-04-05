import { db, user as userTable, eq } from "@repo/db";
import { type listed_job, type customerSignupRequest, type user, type Variables } from "../lib/types/types";
import { Hono } from "hono";
import { requireAuth } from "@repo/auth";
import { fetchHandymen, findRecordsInJobs, listJob } from "../lib/queries";

const customerRouter = new Hono<{Variables: Variables}>();

customerRouter.post('/sign-up', async (c) => {
    const body: customerSignupRequest = await c.req.json();
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

    if (!response.ok) {
        const error = await response.json();
        return c.json({ error: 'Signup failed', details: error }, response.status as 400);
    }

    const data = await response.json() as {token: string, user: user};
    
    if (!data.user || !data.user.id) {
        return c.json({ error: 'Invalid response from auth service' }, 500);
    }

    const id  = data.user.id;
    const token = data.token;

    const [updatedUser] = await db
        .update(userTable)
        .set({ role: 'customer'})
        .where(eq(userTable.id, id))
        .returning()

    return c.json({
        token: data.token,
        user : updatedUser
    });
})

customerRouter.post('/sign-in', async (c) => {
    const { email, password } = await c.req.json();

    const origin = new URL(c.req.url).origin;
    const response = await fetch(`${origin}/auth/sign-in/email`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
        const error = await response.json();
        return c.json({ error: 'Sign in failed', details: error }, response.status as 400);
    }

    const setCookie = response.headers.get('set-cookie');
    if (setCookie) {
        c.header('Set-Cookie', setCookie);
    }

    const data = await response.json();
    return c.json(data);
});

customerRouter.get('/dashboard',requireAuth, async (c) => {
    const user = c.var.user;

    const jobs = await findRecordsInJobs('customer', user.id)

    const completeJobs = jobs.filter(job => job.job_status === "Completed");
    const incompleteJobs = jobs.filter(job => job.job_status === "NotCompleted");

    const completeJobRevenue = completeJobs.reduce((sum,job) => sum + job.cost, 0) 
    const incompleteJobRevenue = incompleteJobs.reduce((sum,job) => sum + job.cost, 0)
    
    const handymen = jobs.map(job => job.handyman)

    c.set("jobs",jobs)

    return c.json({completeJobs, completeJobRevenue, incompleteJobs, incompleteJobRevenue, handymen});
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

customerRouter.post('/direct-hire', async (c) => {
    const handymen = await fetchHandymen()
    const handymenByCategory = handymen.reduce((acc, handyman) => {
        if (!handyman.category) return acc;

        const cat = handyman.category;
        if(!acc[cat]){
            acc[cat] = [];
        }

        acc[cat].push(handyman)

        return acc;
    }, {} as Record<string, typeof handymen>);

    return c.json({
        message: "All avilable handymen",
        handymenByCategory
    }, 200);
})

export default customerRouter;