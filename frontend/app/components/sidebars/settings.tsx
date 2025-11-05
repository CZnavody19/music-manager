import { Link, useLocation } from "react-router"
import { Sidebar, SidebarContent, SidebarFooter, SidebarGroup, SidebarGroupContent, SidebarGroupLabel, SidebarHeader, SidebarMenuButton } from "~/components/ui/sidebar"

const application = [
    { label: "General", to: "/settings" },
]

const sources = [
    { label: "YouTube", to: "/settings/youtube" },
]

const downloaders = [
    { label: "Tidal", to: "/settings/tidal" },
]

const notifications = [
    { label: "Discord", to: "/settings/discord" },
]

const integrations = [
    { label: "Plex", to: "/settings/plex" },
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
                        {application.map((item) => (
                            <SidebarMenuButton asChild isActive={pathname === item.to} key={item.to}>
                                <Link to={item.to} prefetch="intent">{item.label}</Link>
                            </SidebarMenuButton>
                        ))}
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>Sources</SidebarGroupLabel>
                    <SidebarGroupContent>
                        {sources.map((item) => (
                            <SidebarMenuButton asChild isActive={pathname === item.to} key={item.to}>
                                <Link to={item.to} prefetch="intent">{item.label}</Link>
                            </SidebarMenuButton>
                        ))}
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>Downloaders</SidebarGroupLabel>
                    <SidebarGroupContent>
                        {downloaders.map((item) => (
                            <SidebarMenuButton asChild isActive={pathname === item.to} key={item.to}>
                                <Link to={item.to} prefetch="intent">{item.label}</Link>
                            </SidebarMenuButton>
                        ))}
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>Notifications</SidebarGroupLabel>
                    <SidebarGroupContent>
                        {notifications.map((item) => (
                            <SidebarMenuButton asChild isActive={pathname === item.to} key={item.to}>
                                <Link to={item.to} prefetch="intent">{item.label}</Link>
                            </SidebarMenuButton>
                        ))}
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>Integrations</SidebarGroupLabel>
                    <SidebarGroupContent>
                        {integrations.map((item) => (
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