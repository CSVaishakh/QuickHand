  import LandingFooter from "./_components/footer";
  import LandingHeader from "./_components/header"
  import LandingHero from "./_components/hero";

  export default function Home() {
    return(
      <>
        <section className="h-screen flex flex-col bg-blue-400">
          <LandingHeader/>
          <LandingHero/>
          <LandingFooter/>
        </section>
      </>
    )
  }
