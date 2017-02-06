import { EconomyAnalyzerPage } from './app.po';

describe('economy-analyzer App', function() {
  let page: EconomyAnalyzerPage;

  beforeEach(() => {
    page = new EconomyAnalyzerPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
