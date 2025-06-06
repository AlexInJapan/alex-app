// app/login/page.tsx
'use client';

import { useState, FormEvent } from 'react';
import styles from "@/app/page.module.css"
import { useRouter } from 'next/navigation';
import { apiClient } from '@/lib/api';

export default function LoginPage() {
  const [email, setEmail] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [message, setMessage] = useState<string>('');
  const [loginUser, setLoginUser] = useState<string>('');
  const [registerEmail, setRegisterEmail] = useState<string>('');
  const [registerName, setRegisterName] = useState<string>('');
  const [registerLoading, setRegisterLoading] = useState<boolean>(false);
  const [registerMessage, setRegisterMessage] = useState<string>('');
  const router = useRouter();

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);
    setMessage('');

    try {
      const data = await apiClient.login(email);
      if (data) {
        setMessage(data.message || 'リクエストが成功しました！');
        setLoginUser(data.user || '');
        setTimeout(() => {
          router.push('/dashboard');
        }, 2000); // 少し待ってからメッセージを消す
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

  const handleRegister = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setRegisterLoading(true);
    setRegisterMessage('');

    try {
      const data = await apiClient.register(registerEmail, registerName);
      if (data) {
        setRegisterMessage(data.message || '登録が成功しました！');
      } else {
        setRegisterMessage((data && data.message) || '登録に失敗しました。');
      }
    } catch (error) {
      console.error('APIリクエストエラー:', error);
      setRegisterMessage('エラーが発生しました。通信状態を確認するか、時間をおいて再度お試しください。');
    } finally {
      setRegisterLoading(false);
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
      {message && <p className={styles.message}>{message}:{loginUser}</p>}

      {/* Register Form */}
      <form onSubmit={handleRegister} className={styles.form} style={{ marginTop: '2rem' }}>
        <div className={styles.inputGroup}>
          <label htmlFor="registerEmail" className={styles.label}>
            新規登録メールアドレス
          </label>
          <input
            type="email"
            id="registerEmail"
            name="registerEmail"
            value={registerEmail}
            onChange={(e) => setRegisterEmail(e.target.value)}
            required
            className={styles.input}
            placeholder="email@example.com"
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor="registerPassword" className={styles.label}>
            パスワード
          </label>
          <input
            type="text"
            id="name"
            name="name"
            value={registerName}
            onChange={(e) => setRegisterName(e.target.value)}
            required
            className={styles.input}
            placeholder="Your name"
          />
        </div>
        <button type="submit" disabled={registerLoading} className={styles.button}>
          {registerLoading ? '登録中...' : '新規登録'}
        </button>
      </form>
      {registerMessage && <p className={styles.message}>{registerMessage}</p>}
    </div>
  );
}