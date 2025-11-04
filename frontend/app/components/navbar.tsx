import { Link } from "react-router"
import { NavigationMenu, NavigationMenuContent, NavigationMenuIndicator, NavigationMenuItem, NavigationMenuLink, NavigationMenuList, NavigationMenuTrigger, NavigationMenuViewport } from "~/components/ui/navigation-menu"

export function Navbar() {
    return (
        <NavigationMenu className="mx-auto p-2">
            <NavigationMenuList>
                <NavigationMenuItem>
                    <NavigationMenuLink asChild>
                        <Link to="/" prefetch="intent">Dashboard</Link>
                    </NavigationMenuLink>
                </NavigationMenuItem>
                <NavigationMenuItem>
                    <NavigationMenuLink asChild>
                        <Link to="/settings" prefetch="intent">Settings</Link>
                    </NavigationMenuLink>
                </NavigationMenuItem>
                {/* <NavigationMenuItem>
                    <NavigationMenuTrigger>Item One</NavigationMenuTrigger>
                    <NavigationMenuContent>
                        <NavigationMenuLink>Link</NavigationMenuLink>
                    </NavigationMenuContent>
                </NavigationMenuItem> */}
            </NavigationMenuList>
        </NavigationMenu>
    )
}