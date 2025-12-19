export const getJobCategory = (handymanCategory: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "havc_technician" | "landscaper" | "cleaner" ) => {
    const categoryMap = {
        plumber: "plumbing",
        electrician: "electrical",
        carpenter: "carpentery",
        mason: "masonary",
        mechanic: "mechanical",
        havc_technician: "havc",
        landscaper: "landscaping",
        cleaner: "cleaning"
    } as const;

    return categoryMap[handymanCategory]
}