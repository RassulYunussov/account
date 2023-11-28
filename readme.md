# Account library

Helps generating accounts using [IBAN] (https://en.wikipedia.org/wiki/International_Bank_Account_Number) approach with checksum

Initial design - 20 alpha-numerical representation:

- Any first to letters - interpreated as country
- 2 digits - control sum
- Any two letters - interpreated as organization
- "AC" - identifying account
- 12 digits - account number from 1 to 999_999_999_999

Examples:
- KZ98TPAC000007399051