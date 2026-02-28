# WatchTower

**Learning SIEM Concepts by Building a Mini Tool (AI-assisted)**

WatchTower is a small project I built to understand **how SIEM (Security Information and Event Management) systems work**.  
It collects login events and detects suspicious activity like repeated failed login attempts.  

> Built with **AI guidance**, the goal was to learn SIEM concepts, not to create a production-ready system.

---

## What It Does

- Collects login events via REST API (`/api/logs`)  
- Detects repeated failed login attempts  
- Prints alerts in real time for suspicious activity  
- Helps understand **SIEM concepts**: log collection, pattern detection, alerting  

> The backend is written in **Go** with a PostgreSQL database, but the tech is secondary — the focus is **on learning security monitoring**.

---

## Quick Start

1. **Clone the repository:**

```bash
git clone https://github.com/nouraiztee/watchtower.git
cd watchtower