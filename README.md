# mcpdump

[![test](https://github.com/baldaworks/mcpdump/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/baldaworks/mcpdump/actions/workflows/test.yml)
[![lint](https://github.com/baldaworks/mcpdump/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/baldaworks/mcpdump/actions/workflows/lint.yml)
[![security](https://github.com/baldaworks/mcpdump/actions/workflows/security.yml/badge.svg?branch=main)](https://github.com/baldaworks/mcpdump/actions/workflows/security.yml)
[![release](https://github.com/baldaworks/mcpdump/actions/workflows/omnidist-release.yml/badge.svg)](https://github.com/baldaworks/mcpdump/actions/workflows/omnidist-release.yml)
[![npm](https://img.shields.io/npm/v/@baldaworks/mcpdump)](https://www.npmjs.com/package/@baldaworks/mcpdump)
[![License](https://img.shields.io/github/license/baldaworks/mcpdump)](LICENSE)
[![Version](https://img.shields.io/github/v/tag/baldaworks/mcpdump?label=version)](https://github.com/baldaworks/mcpdump/tags)

**Inspect MCP servers before connecting tools or agents.**

`mcpdump` starts any stdio Model Context Protocol (MCP) server command and
prints initialize, capability, tool, prompt, resource, and resource-template
metadata.

Use it when you need to confirm what an MCP server actually exposes before
adding it to an agent, IDE, gateway, or automation workflow.

## What You Get

| Capability | Behavior |
| --- | --- |
| MCP preflight | Starts a stdio MCP server command and initializes a client session. |
| Capability inventory | Prints protocol version, server info, and server capabilities. |
| Tool discovery | Lists tools with input parameters and response schemas. |
| Prompt/resource discovery | Shows prompts, resources, and resource templates when supported. |
| JSON output | Emits machine-readable output for scripts, CI checks, and debugging artifacts. |
| Quiet by default | Keeps inspector lifecycle logs out of the result unless `--debug` is enabled. |

## Try It

Install globally:

```sh
npm install -g @baldaworks/mcpdump@latest
```

Run once with `npx`:

```sh
npx @baldaworks/mcpdump@latest -- <mcp-server-cmd> [args...]
```

Inspect an MCP server:

```sh
mcpdump -- codex mcp-server
mcpdump --json -- codex mcp-server --sandbox workspace-write
mcpdump --debug -- <mcp-server-cmd> [args...]
```

## Provider Commands

| Provider | Command |
| --- | --- |
| Codex MCP server | `mcpdump -- codex mcp-server` |
| Codex with sandbox | `mcpdump --json -- codex mcp-server --sandbox workspace-write` |
| Generic MCP | `mcpdump -- <mcp-server-cmd> [args...]` |

The `--` separator is required. Arguments before `--` are treated as
`mcpdump` flags; arguments after it are passed to the MCP server command.

## Output

Human-readable output includes:

- server name and version
- protocol version
- server capabilities
- tools with input parameters and response schemas
- prompts, resources, and resource templates status

Use `--json` when you want stable output for scripts, test fixtures, issue
reports, or CI logs.

## Use Cases

- Verify an MCP server before registering it with an agent runtime.
- Compare tool schemas across local and remote server commands.
- Capture a compact diagnostic artifact for bug reports.
- Check whether a local command is actually speaking MCP over stdio.

## Flags

| Flag | Purpose |
| --- | --- |
| `--json` | Print machine-readable JSON output. |
| `--debug` | Enable debug logs for the inspector. |
| `-h`, `--help` | Show command help. |

## Repository

- GitHub: <https://github.com/baldaworks/mcpdump>

## Contact

- Issues: <https://github.com/baldaworks/mcpdump/issues>
- Maintainer: [@metalagman](https://github.com/metalagman)

## License

MIT. See the repository [LICENSE](https://github.com/baldaworks/mcpdump/blob/main/LICENSE).
