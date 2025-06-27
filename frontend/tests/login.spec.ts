import { test, expect } from '@playwright/test';

test('login and register flow', async ({ page }) => {
  // Go to the login page
  await page.goto('http://localhost:3000');
  await page.pause();

  // --- Login Form ---
  await page.fill('input[name="email"]', 'testuser@example.com');
  await page.click('button[type="submit"]');

  // Wait for login message (adjust selector if needed)
  await expect(page.locator('p[class*="message"]')).toBeVisible();

  // --- Register Form ---
  await page.fill('input[name="registerEmail"]', 'newuser@example.com');
  await page.fill('input[name="name"]', 'New User');
  await page.click('form:nth-of-type(2) button[type="submit"]');

  // Wait for registration message (adjust selector if needed)
  await expect(page.locator('p[class*="message"]')).toBeVisible();
});
