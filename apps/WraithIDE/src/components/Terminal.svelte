<script>
	import { Xterm, XtermAddon } from '@battlefieldduck/xterm-svelte';
	let terminal;
	let currentLine = '';
	let prompt = 'user@wraith:~/project$ ';
	let options = {
		fontFamily: 'Fira Code, monospace',
		fontSize: 14,
		cursorBlink: true
	};

	async function onLoad() {
		const fitAddon = new (await XtermAddon.FitAddon()).FitAddon();
		terminal.loadAddon(fitAddon);
		fitAddon.fit();
		writePrompt();
	}

	function writePrompt() {
		terminal.write(`\r\n${prompt}`);
		currentLine = '';
	}

	function onKey({ key, domEvent }) {
		const code = domEvent.code;

		if (code === 'Enter') {
			terminal.write('\r\n');
			handleCommand(currentLine.trim());
			writePrompt();
		} else if (code === 'Backspace') {
			if (currentLine.length > 0) {
				currentLine = currentLine.slice(0, -1);
				terminal.write('\b \b');
			}
		} else if (code.startsWith('Arrow')) {
			// Optional: add arrow key history navigation here
			return;
		} else {
			currentLine += key;
			terminal.write(key);
		}
	}

	function handleCommand(cmd) {
		if (cmd === '') return;
		if (cmd === 'help') {
			terminal.write('Available commands: help, clear, echo <text>');
		} else if (cmd.startsWith('echo ')) {
			terminal.write(cmd.slice(5));
		} else if (cmd === 'clear') {
			terminal.clear();
		} else {
			terminal.write(`Command not found: ${cmd}`);
		}
	}
</script>

<Xterm bind:terminal {options} {onLoad} {onKey} class="h-[30%] border-t text-white" />
