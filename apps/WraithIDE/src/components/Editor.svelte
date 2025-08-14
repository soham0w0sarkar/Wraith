<script>
	import { onDestroy, onMount } from 'svelte';
	import catppuccinTheme from '../theme/catppuccin-mocha.json';

	let editor;
	let monaco;
	let editorContainer;

	onMount(async () => {
		try {
			console.log('Loading Monaco Editor...');
			monaco = (await import('../monaco')).default;

			monaco.editor.defineTheme('catppuccin-mocha', catppuccinTheme);

			// Create the editor with the theme applied
			editor = monaco.editor.create(editorContainer, {
				theme: 'catppuccin-mocha',
				fontSize: 14,
				fontFamily: 'JetBrains Mono, Fira Code, monospace',
				lineHeight: 1.5,
				minimap: { enabled: false },
				scrollBeyondLastLine: false,
				automaticLayout: true,
				tabSize: 2,
				insertSpaces: true,
				wordWrap: 'on',
				lineNumbers: 'on',
				renderWhitespace: 'selection',
				bracketPairColorization: {
					enabled: true
				}
			});

			// Create and set the model
			const model = monaco.editor.createModel(
				`// Welcome to Monaco Editor with Catppuccin Mocha theme!
console.log('Hello from Monaco! (the editor, not the city...)');

const greet = (name) => {
	return \`Hello, \${name}! 🎉\`;
};

const numbers = [1, 2, 3, 4, 5];
const doubled = numbers.map(n => n * 2);

// Try some syntax highlighting
class CatppuccinDemo {
	constructor(theme) {
		this.theme = theme;
		this.isAwesome = true;
	}
	
	getThemeInfo() {
		return {
			name: this.theme,
			colors: ['#1e1e2e', '#cba6f7', '#a6e3a1', '#f9e2af'],
			mood: 'cozy' 
		};
	}
}

// RegExp example
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

// Some constants
const PI = 3.14159;
const greeting = "Hello, world!";
const isValid = true;`,
				'javascript'
			);

			editor.setModel(model);

			console.log('Monaco Editor loaded successfully with Catppuccin theme!');
		} catch (error) {
			console.error('Error loading Monaco Editor:', error);
		}
	});

	onDestroy(() => {
		monaco?.editor.getModels().forEach((model) => model.dispose());
		editor?.dispose();
	});
</script>

<div class="preset-glass-neutral flex h-[70%] w-full flex-col">
	<span class="border-primary-300 block border-b p-1 font-mono text-center">Wraith.shh</span>
	<div class="h-screen w-full" bind:this={editorContainer}></div>
</div>

<style></style>
