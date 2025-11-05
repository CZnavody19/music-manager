import { Music4 } from "lucide-react"
import { Link } from "react-router"
import { Button } from "~/components/ui/button"
import { NavigationMenu, NavigationMenuItem, NavigationMenuLink, NavigationMenuList } from "~/components/ui/navigation-menu"

export function Navbar() {
    return (
        <div className="flex flex-row w-full justify-between items-center py-3 px-6">
            <Music4 size={32} className="text-primary" />
            <NavigationMenu>
                <NavigationMenuList>
                    <NavigationMenuItem>
                        <NavigationMenuLink asChild>
                            <Link to="/" prefetch="intent">Dashboard</Link>
                        </NavigationMenuLink>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <NavigationMenuLink asChild>
                            <Link to="/services" prefetch="intent">Services</Link>
                        </NavigationMenuLink>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <NavigationMenuLink asChild>
                            <Link to="/settings" prefetch="intent">Settings</Link>
                        </NavigationMenuLink>
                    </NavigationMenuItem>
                </NavigationMenuList>
            </NavigationMenu>
            <Button asChild>
                <Link to="/logout" prefetch="intent">Logout</Link>
            </Button>
        </div>
    )
}