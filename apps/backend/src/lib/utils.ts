export const getJobCategory = (handymanCategory: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "havc_technician" | "landscaper" | "deep_cleaner" ) => {
    const categoryMap = {
        plumber: "plumbing",
        electrician: "electrical",
        carpenter: "carpentery",
        mason: "masonary",
        mechanic: "mechanical",
        havc_technician: "havc",
        landscaper: "landscaping",
        deep_cleaner: "deep_cleaning"
    } as const;

    return categoryMap[handymanCategory]
}