import { test, expect } from '@playwright/test';
const URL = 'http://localhost:4500/'

test('Should show the unauthorized forced area', async ({ page }) => {
  await page.goto(URL);

  // Expect a title "to contain" a substring.
  let locator = page.locator(".unauthorized-forced-area")
  let isVisible = await locator.isVisible()
  expect(isVisible).toBe(true)

  locator = page.locator(".btn-secondary")
  isVisible = await locator.isVisible()
  expect(isVisible).toBe(true)
  
  locator.click()
});
 