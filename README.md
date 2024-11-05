# Passive Reconnaissance Tool

## Overview

This project is a passive reconnaissance tool designed to help users become familiar with open-source intelligence (OSINT) techniques for information gathering. Information gathering is a critical and often lengthy step in penetration testing (pentesting), as it lays the groundwork for further security analysis. This tool automates some passive OSINT methods, allowing you to retrieve and organize specific data based on user input.

## Objectives

The primary objective of this project is to help users become comfortable with open-source investigative methods, specifically using a passive recognition tool in Go. The tool gathers data based on various types of input—such as full names, IP addresses, or usernames—and then organizes this data into a report.

## Features

The tool can process the following types of input:
- **Full Name**: Takes "Last name" and "First name" as input, searches for relevant details, and retrieves associated telephone numbers and addresses.
- **IP Address**: Finds at least the city and the Internet Service Provider (ISP) based on the IP address.
- **Username**: Checks if the specified username exists on at least 5 well-known social networks.

The results are saved in a `result.txt` file (or `result2.txt` if `result.txt` already exists).


## Usage

This task is written in Go. You can run project using these commands:
```
go build passive
./passive -fn "Jean Dupont"
./passive -ip 176.112.156.215
./passive -u "@user01"
```
or by running script file:
```
sh run.sh
```

## Author

[Viktoriia/vavstanc](https://01.kood.tech/git/vavstanc)