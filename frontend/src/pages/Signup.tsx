import { useState } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "../components/ui/card"
import { Button } from "../components/ui/button"
import { Input } from "../components/ui/input"
import { Label } from "../components/ui/label"
import { Link, useNavigate } from "react-router-dom";
//import { useAuthStore } from "../store/authstore"

export default function Signup() {
    const [email, setEmail] = useState("")
    const [username, setusername] = useState("")
    const [firstname, setfirstname] = useState("")
    const [lastname, setlastname] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("")
    const [loading, setLoading] = useState(false)
    const navigate = useNavigate();
    //const signup = useAuthStore((state) => state.signup)

    async function handleSubmit(e: React.FormEvent) {



        e.preventDefault()
        setLoading(true)
        setError("")

        try {
            const res = await fetch(`${import.meta.env.VITE_URL}/api/signup`, {

                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, username, firstname, lastname, password, }),
            })

            const data = await res.json()

            if (!res.ok) {
                console.error("Backend error:", data)
                throw new Error(data.error || "Signup failed")
            }

            alert("Signup successful")
            navigate("/", { replace: true });

        }
        catch (err: unknown) {
            console.error("Frontend caught error:", err)

            if (err instanceof Error) {
                setError(err.message)
            } else {
                setError("Something went wrong")
            }
        }
        finally {
            setLoading(false)
        }
    }

    return (
        <div className="flex min-h-screen items-center justify-center bg-background">
            <Card className="w-full max-w-sm">
                <CardHeader>
                    <CardTitle className="text-center text-2xl">Signup</CardTitle>
                </CardHeader>

                <CardContent>
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                            <Label htmlFor="username">User Name</Label>
                            <Input id="username" type="username" value={username} onChange={(e) => setusername(e.target.value)} required
                            />
                        </div>
                        <div>
                            <Label htmlFor="email">Email</Label>
                            <Input
                                id="email"
                                type="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                            />
                        </div>
                        <div>
                            <Label htmlFor="firstname">First Name</Label>
                            <Input id="firstname" type="firstname" value={firstname} onChange={(e) => setfirstname(e.target.value)} required
                            />
                        </div>
                        <div>
                            <Label htmlFor="lastname">Last Name</Label>
                            <Input id="lastname" type="lastname" value={lastname} onChange={(e) => setlastname(e.target.value)} required
                            />
                        </div>
                        <div>
                            <Label htmlFor="password">Password</Label>
                            <Input id="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required
                            />
                        </div>

                        {error && (
                            <p className="text-sm text-red-500">{error}</p>
                        )}

                        <Button className="w-full" disabled={loading}>
                            {loading ? "Logging in..." : "Login"}
                        </Button>
                    </form>
                </CardContent>
                <ul>
                    <li>
                        <Link to="/">Use Signed Out</Link>
                    </li>
                </ul>

            </Card>
        </div>
    )
}