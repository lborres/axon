import { Button } from "@axon/ui/button";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: App,
});

function App() {
  return (
    <>
      <Button variant={"outline"}>Test</Button>
    </>
  );
}
