import { NextResponse } from "next/server"

export async function POST(request: Request) {
    try {
        const body = await request.json();
        const email = body.email;
        const name = body.name;

        if (!email || email.trim() === "") {
            return NextResponse.json({ error: "Enter your email" }, { status: 400 });
        }

        if (!name || name.trim() === "") {
            return NextResponse.json({ error: "Enter your name" }, { status: 400 });
        }


        const res = await fetch("http://backend:8080/createUser", {
            method: "POST",
            headers: {
            "Content-Type": "application/json"
            },
            body: JSON.stringify({ email, name })
        });

        if (!res.ok) {
            return NextResponse.json({ error: "Failed to fetch user data" }, { status: res.status });
        }

        const user = await res.json();
        return NextResponse.json(user, { status: 200 });
    } catch (error) {
        return NextResponse.json({ error: error || "Unknown error" }, { status: 500 });
    }
}