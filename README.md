# Resource library for golang

This repository provides generic interfaces for dealing with repositories of resources.
The two main components are repositories (Repository) and adapters to resources (Adapter).

Currently only the implemenation of FileRepository and FileAdapter are provided.

Other repository types could be a container repository which handles multiple repositories or a network adapter in order to access resources oer a network connection (e.g. TCP).
