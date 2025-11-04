import { Link, useLocation } from "react-router"
import { Sidebar, SidebarContent, SidebarFooter, SidebarGroup, SidebarGroupContent, SidebarGroupLabel, SidebarHeader, SidebarMenuButton } from "~/components/ui/sidebar"

const items = [
    { label: "General", to: "/settings" },
    { label: "YouTube", to: "/settings/youtube" },
    { label: "Notifications", to: "/settings/notifications" },
]

export function SettingsSidebar() {
    const { pathname } = useLocation();

    return (
        <Sidebar collapsible="none" className="bg-background">
            <SidebarHeader>
                <h1 className="font-semibold text-xl px-2">Settings</h1>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>Application</SidebarGroupLabel>
                    <SidebarGroupContent>
                        {items.map((item) => (
                            <SidebarMenuButton asChild isActive={pathname === item.to} key={item.to}>
                                <Link to={item.to} prefetch="intent">{item.label}</Link>
                            </SidebarMenuButton>
                        ))}
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter>
            </SidebarFooter>
        </Sidebar>
    )
}