// app/login/page.tsx
'use client';

import { useState, FormEvent } from 'react';
import styles from "@/app/page.module.css"
import { useRouter } from 'next/navigation'; // 必要に応じて

export default function LoginPage() {
  const [email, setEmail] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [message, setMessage] = useState<string>('');
  const router = useRouter(); // 必要に応じて

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);
    setMessage('');

    const apiPath = '/api/login'

    try {
      const response = await fetch(apiPath, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: email }),
      });
      const data = await response.json(); // レスポンスボディをJSONとしてパース

      if (response.ok) {
        setMessage(data.message || 'リクエストが成功しました！');
        // 例: ログイン成功後にダッシュボードへリダイレクト
        router.push('/dashboard');
      } else {
        setMessage(data.message || 'リクエストに失敗しました。');
      }
    } catch (error) {
      console.error('APIリクエストエラー:', error);
      setMessage('エラーが発生しました。通信状態を確認するか、時間をおいて再度お試しください。');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Eメール入力</h1>
      <form onSubmit={handleSubmit} className={styles.form}>
        <div className={styles.inputGroup}>
          <label htmlFor="email" className={styles.label}>
            メールアドレス
          </label>
          <input
            type="email"
            id="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            className={styles.input}
            placeholder="email@example.com"
          />
        </div>
        <button type="submit" disabled={isLoading} className={styles.button}>
          {isLoading ? '送信中...' : '送信'}
        </button>
      </form>
      {message && <p className={styles.message}>{message}</p>}
    </div>
  );
}