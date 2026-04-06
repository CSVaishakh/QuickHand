export type headerProps = {
    buttons: Array<{
        buttonName: string,
        url?: string,
        onClick?: () => void
    }>;
}

export type footerProps = {
    buttons: Array<{
        buttonName: string,
        url?: string,
        onClick?: () => void
    }>;
}