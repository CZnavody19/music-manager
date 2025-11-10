import { gql } from "@apollo/client";
import { useSubscription } from "@apollo/client/react";
import { Check, Music4, RefreshCw } from "lucide-react"
import { useEffect, useState } from "react";
import { Link, useRevalidator } from "react-router"
import { Button } from "~/components/ui/button"
import { NavigationMenu, NavigationMenuItem, NavigationMenuLink, NavigationMenuList } from "~/components/ui/navigation-menu"
import type { Subscription, Task } from "~/graphql/gen/graphql";

export function Navbar() {
    const { revalidate } = useRevalidator();
    const [task, setTask] = useState<Task | null>(null);
    const [startTime, setStartTime] = useState<number>(0);
    const [time, setTime] = useState<number>(0);
    useSubscription<Subscription>(gql`
        subscription tasks {
            tasks {
                title
                startedAt
                ended
            }
        }
    `, {
        onData: ({ data }) => {
            setTask(data.data?.tasks || null);
            if (data.data?.tasks && !data.data?.tasks.ended) {
                setStartTime(new Date(data.data.tasks.startedAt).getTime());
            } else {
                revalidate();
                setStartTime(0);
                setTimeout(() => {
                    setTime(0);
                    setTask(null);
                }, 5000);
            }
        },
    });

    useEffect(() => {
        let interval: NodeJS.Timeout;
        if (startTime !== 0) {
            interval = setInterval(() => {
                setTime(Math.floor((Date.now() - startTime)));
            }, 100);
        }
        return () => {
            if (interval) {
                clearInterval(interval);
            }
        }
    }, [startTime]);

    return (
        <div className="flex flex-row w-full justify-between items-center py-3 px-6 relative">
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
                            <Link to="/sources" prefetch="intent">Sources</Link>
                        </NavigationMenuLink>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <NavigationMenuLink asChild>
                            <Link to="/integrations" prefetch="intent">Integrations</Link>
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
            {task && (
                <div className="flex flex-row items-center border rounded-md px-2 py-1 gap-2 absolute top-full right-0 mr-4">
                    {task.ended ? <Check /> : <RefreshCw className="animate-spin" />}
                    <div className="flex flex-col">
                        <h5 className="text-sm">{task.title}</h5>
                        <p className="text-xs text-muted-foreground">{time} ms</p>
                    </div>
                </div>
            )}
        </div>
    )
}