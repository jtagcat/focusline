# focusline
```sh
go install github.com/jtagcat/focusline@latest
```

Focuses text to nth character. Dependency for cowsay, but for other animals

Given a wrapping point, say 80 chars, it will shift the text from focus to fit it.
When needed, it will wrap lines.

```ascii
┌──────────────────────────┐ ┌──────────────────────────┐
│Centers text around focus │ │ Shifts to left on EOLine │
└┬────────────────────────┬┘ └┬────────────────────────┬┘
 │               ↓        │   │               ↓        │
 │       lorem ipsum dolor│   │   lorem ipsum dolor sit│
 │               /        │   │               /        │
 │              /         │   │              /         │
 │^__^         /          │   │             /          │
 │(oo)\_______            │   └────────────────────────┘
 │(__)\       )\/\        │    ┌───────────────────────┐
 │    ||----w | (cow from │    │    focus direction    │
 │    ||     ||    cowsay)│    └─┬─────┬───────┬─────┬─┘
 └────────────────────────┘      │  ↓  │   ↓   │  ↓  │
  ┌───────────────────────┐      │focus│ focus │focus│
  │focus last, align rest │      │  on │  and  │ on  │
  └──┬─────────────────┬──┘      │ left│breathe│rigt │
     │         ↓       │         │  /  │  /|\  │  \  │
     │iseenda eest on  │         │ /   │ / | \ │   \ │
     │kõige raskem     │         └─────┴───────┴─────┘
     │     põgeneda    │
     │         \       │
     │          \ left │      Inspired by    Drawn in
     ├─────────────────┤             cowsay    asciiflow
     │           ↓     │
     │  iseenda eest on│          │ focusline      │
     │     kõige raskem│          │     by jtagcat │
     │        põgeneda │
     │           /     │
     │    right /      │     Examples in README are tested.
     └─────────────────┘
```
