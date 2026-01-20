test('should register and login, check success message and redirect', async ({ page }) => {
  await page.goto('http://localhost:3000');

  // Generate random credentials
  const randomEmail = `user${Math.floor(Math.random() * 10000)}@example.com`;
  const randomPassword = `pw${Math.floor(Math.random() * 10000)}`;

  // Register the account
  await page.getByLabel('新規登録メールアドレス', { exact: true }).fill(randomEmail);
  await page.getByPlaceholder('Your name').fill(randomPassword);
  await page.getByRole('button', { name: '新規登録' }).click();
  // Check for success message in paragraph after registration
  await expect(page.getByText(/User created successfully|登録が成功しました|登録完了|success|registered/i)).toBeVisible();

  // Login with the registered email
  await page.getByLabel('メールアドレス', { exact: true }).fill(randomEmail);
  await page.getByRole('button', { name: '送信' }).click();
  // Check for login success message in paragraph
  await expect(page.getByText(/ログイン成功|ログイン完了|success|logged in/i)).toBeVisible();

  // Check for redirect (URL change or dashboard heading)
  await expect(page).not.toHaveURL('http://localhost:3000');
  // Optionally, check for dashboard heading
  // await expect(page.getByRole('heading', { name: /dashboard|ダッシュボード/i })).toBeVisible();
});
test('should register a random account and login with it', async ({ page }) => {
  await page.goto('http://localhost:3000');

  // Generate random credentials
  const randomEmail = `user${Math.floor(Math.random() * 10000)}@example.com`;
  const randomPassword = `pw${Math.floor(Math.random() * 10000)}`;

  // Register the account
  await page.getByLabel('新規登録メールアドレス', { exact: true }).fill(randomEmail);
  await page.getByPlaceholder('Your name').fill(randomPassword);
  await page.getByRole('button', { name: '新規登録' }).click();
  await expect(page.getByRole('alert')).toBeVisible();

  // Try to login with the registered email
  await page.getByLabel('メールアドレス', { exact: true }).fill(randomEmail);
  await page.getByRole('button', { name: '送信' }).click();
  await expect(page.getByRole('alert')).toBeVisible();
});
import { test, expect } from '@playwright/test';

test('should display email input and registration form', async ({ page }) => {
  await page.goto('http://localhost:3000');

  // Check for main heading
  await expect(page.getByRole('heading', { name: 'Eメール入力' })).toBeVisible();

  // Check for email input and send button
  await expect(page.getByLabel('メールアドレス', { exact: true })).toBeVisible();
  await expect(page.getByRole('button', { name: '送信' })).toBeVisible();

  // Check for registration section
  await expect(page.getByLabel('新規登録メールアドレス', { exact: true })).toBeVisible();
  await expect(page.getByPlaceholder('Your name')).toBeVisible();
  await expect(page.getByRole('button', { name: '新規登録' })).toBeVisible();
});

test('should submit email and registration forms with random input', async ({ page }) => {
  await page.goto('http://localhost:3000');

  // Fill and submit email input
  const randomEmail = `user${Math.floor(Math.random() * 10000)}@example.com`;
  await page.getByLabel('メールアドレス', { exact: true }).fill(randomEmail);
  await page.getByRole('button', { name: '送信' }).click();

  // Fill and submit registration form
  const randomRegEmail = `reg${Math.floor(Math.random() * 10000)}@example.com`;
  const randomPassword = `pw${Math.floor(Math.random() * 10000)}`;
  await page.getByLabel('新規登録メールアドレス', { exact: true }).fill(randomRegEmail);
  await page.getByPlaceholder('Your name').fill(randomPassword);
  await page.getByRole('button', { name: '新規登録' }).click();

  // Check for a success message or alert after submission
  await expect(page.getByRole('alert')).toBeVisible();
});
