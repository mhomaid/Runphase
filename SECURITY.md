# Security Policy

## Supported versions

Runphase is pre-alpha. Only the `main` branch receives fixes.

## Reporting a vulnerability

Please do not open a public issue.

Use GitHub's private vulnerability reporting on this repository (Security → Report a vulnerability). Include:
- affected component and version or commit;
- steps to reproduce;
- impact.

You should receive an acknowledgement within 5 business days.

## Scope notes

- Runbooks are typed data; arbitrary code execution through runbooks is a vulnerability.
- Runphase must not log prompts, secrets, or credentials by default.
- External endpoint URLs are validated; SSRF through target registration is in scope.
