import { Outlet } from "react-router";
import { SettingsSidebar } from "~/components/sidebars/settings";
import { SidebarProvider } from "~/components/ui/sidebar";

export default function Layout() {
    return (
        <div className="flex flex-row w-full h-full">
            <SidebarProvider>
                <SettingsSidebar />
                <div className="flex flex-col w-full h-full">
                    <Outlet />
                </div>
            </SidebarProvider>
        </div>
    )
}