"use client";
import { Suspense } from "react";
import { useSearchParams } from "next/navigation";
import styles from "@/app/dashboard/page.module.css";

function DashboardContent() {
  const searchParams = useSearchParams();
  const userName = searchParams.get('username');
  return (
    <main className={styles.dashboardMain}>
      <header className={styles.dashboardHeader}>
        <h1 className={styles.dashboardTitle}>Hello, {userName}!</h1>
        <p className={styles.dashboardSubtitle}>Welcome to your dashboard</p>
      </header>
      <section>
        <h2 className={styles.dashboardSectionTitle}>Quick Links</h2>
        <ul className={styles.dashboardList}>
          <li className={styles.dashboardListItem}>Profile</li>
          <li className={styles.dashboardListItem}>Settings</li>
          <li className={styles.dashboardListItem}>Recent Activity</li>
        </ul>
      </section>
    </main>
  );
}

export default function Dashboard() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <DashboardContent />
    </Suspense>
  );
}